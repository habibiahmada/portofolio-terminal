// Package main is the entry point for the SSH experience.
// It runs a Wish SSH server that serves the TUI to remote users.
//
// Security model: this server is intentionally public (portfolio showcase).
// No client authentication is required — visitors connect and browse the TUI.
// Hardening applied:
//   - Remote command execution blocked (interactive PTY sessions only)
//   - Per-IP connection rate limiting and global concurrent session cap
//   - Idle and absolute session timeouts
//   - Host key persisted at SSH_HOST_KEY_PATH (default .ssh/term_info_ed25519)
//   - Port forwarding denied by default (charmbracelet/ssh default)
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync/atomic"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/accesscontrol"
	"github.com/charmbracelet/wish/activeterm"
	bm "github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
	"github.com/charmbracelet/wish/ratelimiter"
	"golang.org/x/time/rate"

	"github.com/habibiahmada/habibiahmada-terminal/internal/tui"
)

const (
	defaultHost        = "0.0.0.0"
	defaultPort        = 2222
	defaultIdleTimeout = 30 * time.Minute
	defaultMaxTimeout  = 2 * time.Hour
	defaultMaxSessions = 64
	defaultRateLimit   = 1 // connections per second per IP
	defaultRateBurst   = 5
	defaultRateEntries = 1024
)

func main() {
	host := envString("SSH_HOST", defaultHost)
	port := envInt("SSH_PORT", defaultPort)
	idleTimeout := envDuration("SSH_IDLE_TIMEOUT", defaultIdleTimeout)
	maxTimeout := envDuration("SSH_MAX_TIMEOUT", defaultMaxTimeout)
	maxSessions := envInt("SSH_MAX_SESSIONS", defaultMaxSessions)
	rateLimit := envFloat("SSH_RATE_LIMIT", defaultRateLimit)
	rateBurst := envInt("SSH_RATE_BURST", defaultRateBurst)

	hostKeyPath := os.Getenv("SSH_HOST_KEY_PATH")
	if hostKeyPath == "" {
		hostKeyPath = ".ssh/term_info_ed25519"
	}

	srv, err := wish.NewServer(
		wish.WithAddress(fmt.Sprintf("%s:%d", host, port)),
		wish.WithHostKeyPath(hostKeyPath),
		wish.WithIdleTimeout(idleTimeout),
		wish.WithMaxTimeout(maxTimeout),
		wish.WithMiddleware(
			bm.Middleware(teaHandler),
			logging.Middleware(),
			maxSessionsMiddleware(maxSessions),
			activeterm.Middleware(),
			accesscontrol.Middleware(), // blocks ssh host <cmd>; interactive shell OK
			ratelimiter.Middleware(ratelimiter.NewRateLimiter(
				rate.Limit(rateLimit),
				rateBurst,
				defaultRateEntries,
			)),
		),
	)
	if err != nil {
		log.Fatalf("Could not create server: %v", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	log.Printf("Starting public SSH portfolio on %s:%d (max_sessions=%d idle=%s max=%s)",
		host, port, maxSessions, idleTimeout, maxTimeout)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("Could not start server: %v", err)
		}
	}()

	<-done
	log.Println("Shutting down SSH server...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Could not shutdown server: %v", err)
	}
}

// teaHandler creates a new Bubble Tea model for each SSH session.
func teaHandler(sess ssh.Session) (tea.Model, []tea.ProgramOption) {
	model := tui.NewSplash()

	return model, []tea.ProgramOption{
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	}
}

func maxSessionsMiddleware(max int) wish.Middleware {
	var active int32
	return func(next ssh.Handler) ssh.Handler {
		return func(sess ssh.Session) {
			if max <= 0 {
				next(sess)
				return
			}
			n := atomic.AddInt32(&active, 1)
			if int(n) > max {
				atomic.AddInt32(&active, -1)
				wish.Fatalln(sess, "Server busy — too many active sessions. Try again later.")
				return
			}
			defer atomic.AddInt32(&active, -1)
			next(sess)
		}
	}
}

func envString(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func envInt(key string, fallback int) int {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		log.Printf("invalid %s=%q, using default %d", key, v, fallback)
		return fallback
	}
	return n
}

func envFloat(key string, fallback float64) float64 {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	n, err := strconv.ParseFloat(v, 64)
	if err != nil {
		log.Printf("invalid %s=%q, using default %v", key, v, fallback)
		return fallback
	}
	return n
}

func envDuration(key string, fallback time.Duration) time.Duration {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		log.Printf("invalid %s=%q, using default %s", key, v, fallback)
		return fallback
	}
	return d
}
