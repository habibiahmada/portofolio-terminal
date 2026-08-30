// Package main is the entry point for the SSH experience.
// It runs a Wish SSH server that serves the TUI to remote users.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	bm "github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"

	"github.com/habibiahmada/habibiahmada-terminal/internal/tui"
)

const (
	host = "0.0.0.0"
	port = 2222
)

func main() {
	// Build the Wish SSH server.
	srv, err := wish.NewServer(
		wish.WithAddress(fmt.Sprintf("%s:%d", host, port)),
		wish.WithHostKeyPath(".ssh/term_info_ed25519"),
		wish.WithMiddleware(
			bm.Middleware(teaHandler),
			logging.Middleware(),
		),
	)
	if err != nil {
		log.Fatalf("Could not create server: %v", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	log.Printf("Starting SSH server on %s:%d", host, port)
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
	// Fresh splash per session; it transitions into the shared TUI core.
	model := tui.NewSplash()

	return model, []tea.ProgramOption{
		tea.WithAltScreen(),
	}
}
