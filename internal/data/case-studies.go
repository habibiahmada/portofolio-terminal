package data

// GetCaseStudy returns the case study for a project slug (live API or bundled).
func GetCaseStudy(slug string) *CaseStudy {
	if cs, ok := liveCaseStudy(slug); ok {
		return cs
	}
	return bundledCaseStudy(slug)
}

func bundledCaseStudy(slug string) *CaseStudy {
	for i := range bundledCaseStudies {
		if bundledCaseStudies[i].Slug == slug {
			cs := bundledCaseStudies[i]
			return &cs
		}
	}
	return nil
}

// GetCaseStudies returns all bundled case studies (offline catalog).
func GetCaseStudies() []CaseStudy {
	return bundledCaseStudies
}

var bundledCaseStudies = []CaseStudy{
	{
		Slug:  "aksara-pustaka",
		Title: "Aksara Pustaka",
		Year:  "2026",
		Tags:  []string{"Laravel", "MySQL", "PHP", "Tailwind CSS"},
		Hero:  "A web library system that brought books, members, loans, returns, stock, and history into one place.",
		Sections: []CaseStudySection{
			{Label: "Where it started", Body: "The school library kept records in notebooks and loose sheets. Looking up whether a book was available, who had it, or when it was due meant thumbing through pages. Aksara Pustaka started as a way to put every one of those questions behind a search box."},
			{Label: "What boxed the work in", Body: "The system had to be simple enough for staff with varied computer comfort to use every day, and it had to keep working on the modest hardware the school already had. No cloud dependency, no always-on external service."},
			{Label: "How it came together", Body: "I mapped the core flows — cataloguing, member registration, loan, return, and stock — before writing code. Laravel handled the domain and MySQL the records, with a Tailwind UI that kept forms and tables readable on small screens."},
			{Label: "What I can stand behind", Body: "Every transaction leaves a trace. Stock and loan history are accounted for, so the librarian can answer any question about the collection without opening a notebook."},
		},
	},
	{
		Slug:  "sipadu",
		Title: "SiPadu",
		Year:  "2026",
		Tags:  []string{"Laravel", "Tailwind CSS", "MySQL", "PHP"},
		Hero:  "A real-time facility reporting system that replaced paper trails for school infrastructure.",
		Sections: []CaseStudySection{
			{Label: "Where it started", Body: "Facility complaints — a broken projector, a leaking roof, a faulty socket — were reported verbally or on paper, then forgotten. SiPadu was built to make a report as quick as sending a message and to make its status visible to everyone involved."},
			{Label: "What boxed the work in", Body: "Users report from phones with poor screens and older browsers. The form had to be tiny, the flow forgiving, and the status view obvious — nobody should need training to file a report."},
			{Label: "How it came together", Body: "I focused on a thin vertical slice: report, assign, resolve, track. Laravel provided auth and the data model, while a Tailwind, mobile-first UI kept the interaction simple across departments."},
			{Label: "What I can stand behind", Body: "Reports now have owners and timestamps. Instead of vanishing into paper, a complaint becomes a tracked item anyone on staff can check on."},
		},
	},
	{
		Slug:  "parking-app",
		Title: "ParkingApp",
		Year:  "2026",
		Tags:  []string{"PHP"},
		Hero:  "A parking web app for check-in, check-out, reporting, and audit in one flow.",
		Sections: []CaseStudySection{
			{Label: "Where it started", Body: "Manual parking records made revenue and occupancy impossible to audit — totals lived in memory and scribbled notes. ParkingApp turned check-in, check-out, and reporting into one accountable flow."},
			{Label: "What boxed the work in", Body: "The operator needs to log a vehicle in seconds without dropping the phone or the pen. The stack had to run on cheap shared hosting with nothing beyond PHP available."},
			{Label: "How it came together", Body: "Starting from a single path — check in a vehicle, check it out, see the session total — I added role-based access and daily reporting on top, keeping the operator screen fast and tactile."},
			{Label: "What I can stand behind", Body: "Every parking session is a row of truth: who, when, in, out, and what was paid. Disputes and end-of-day reconciliation are now arithmetic the system does for you."},
		},
	},
	{
		Slug:  "inventoryflow",
		Title: "Inventoryflow",
		Year:  "2026",
		Tags:  []string{"Laravel", "Tailwind CSS", "PHP", "MySQL"},
		Hero:  "School and lab equipment loans, inventory, approvals, and returns — without spreadsheet chaos.",
		Sections: []CaseStudySection{
			{Label: "Where it started", Body: "Lab equipment loans ran through shared spreadsheets that went out of date the moment two people edited them. Inventoryflow replaced the chaos with a single source of truth for what exists and where it is."},
			{Label: "What boxed the work in", Body: "Loans need an approval step, so only authorised staff can sign gear out. The UI had to handle many rows of equipment without turning into a wall of text."},
			{Label: "How it came together", Body: "I built the loan lifecycle first — request, approve, lend, return — then layered inventory counts and audit history. Approvals happen in-app with a clear paper trail."},
			{Label: "What I can stand behind", Body: "The school now knows what it owns, who has it, and when it is due back. Equipment that used to disappear now has an owner and a due date."},
		},
	},
	{
		Slug:  "bagiberkah",
		Title: "BagiBerkah",
		Year:  "2026",
		Tags:  []string{"Next.js", "Express", "Prisma", "Mayar", "Xendit"},
		Hero:  "A digital THR experience with mini-games and smarter allocation recommendations.",
		Sections: []CaseStudySection{
			{Label: "Where it started", Body: "Sharing THR during a holiday is a social moment, but splitting amounts across family, savings, and zakat is fiddly and forgettable. BagiBerkah made it feel like a game while nudging better allocation."},
			{Label: "What boxed the work in", Body: "Payments needed to work across multiple local gateways (Mayar, Xendit), and the flow had to stay light enough to share on a phone over messenger."},
			{Label: "How it came together", Body: "A Next.js front end drove the mini-games and allocation UI, while Express, Prisma, and a relational store handled transactions and accounting. Each gateway was wrapped behind one guarded interface."},
			{Label: "What I can stand behind", Body: "The money side is auditable and the fun side is shareable. Allocation advice is grounded in real splits people actually use, not a gimmick."},
		},
	},
	{
		Slug:  "e-vote",
		Title: "E-Democracy (E-Vote)",
		Year:  "2025",
		Tags:  []string{"PHP", "Laravel", "Bootstrap", "MySQL"},
		Hero:  "Digital OSIS elections for students — from ballots to real-time results with confidence.",
		Sections: []CaseStudySection{
			{Label: "Where it started", Body: "Paper ballots made school elections slow to count and easy to question. E-Vote digitised the process so students could cast a ballot and watch honest, real-time results appear."},
			{Label: "What boxed the work in", Body: "A school election has hard rules: one vote per student, no double voting, and a result everyone trusts. The system had to prove correctness, not just show numbers."},
			{Label: "How it came together", Body: "Laravel enforced the integrity model — verified voters, a single ballot each, auditable tallies — while a Bootstrap UI kept voting clear on shared school devices."},
			{Label: "What I can stand behind", Body: "The count is verifiable and the rules are enforced by the system. Students see a fair result they can trust instead of a suspicious tally."},
		},
	},
	{
		Slug:  "agrify",
		Title: "Smartfarm AI (Agrify)",
		Year:  "2025",
		Tags:  []string{"React", "Python", "Machine Learning", "JavaScript"},
		Hero:  "A team project combining modern web UI with ML-assisted insights for Indonesian farmers.",
		Sections: []CaseStudySection{
			{Label: "Where it started", Body: "Farmers have expert knowledge but little data at their fingertips. Agrify turned model output into advice a farmer can act on, rather than another dashboard only engineers read."},
			{Label: "What boxed the work in", Body: "The audience is not a data scientist — it is a person in a field. Insights had to be plain-language and actionable, and the ML backend had to serve them quickly."},
			{Label: "How it came together", Body: "I owned the product surface: translating model predictions into clear, decision-ready advice behind a React interface, wired to a Python ML service."},
			{Label: "What I can stand behind", Body: "The project won a country award at a global AI festival. My part was making sure the intelligence didn't stay trapped in a notebook — it reached the people it was meant for."},
		},
	},
	{
		Slug:  "culture-connect",
		Title: "CultureConnect.",
		Year:  "2025",
		Tags:  []string{"Python", "React", "Machine Learning", "Express", "Prisma"},
		Hero:  "An AI platform for more personal, community-aware cultural travel, built with a distributed team.",
		Sections: []CaseStudySection{
			{Label: "Where it started", Body: "Generic travel recommendations show tourists the same crowded spots and miss local culture. CultureConnect aimed to recommend experiences that feel personal and community-aware."},
			{Label: "What boxed the work in", Body: "The team was distributed and the deadline real — part of a competitive programme. Coordinating a ML model, an API, and a React client across time zones took deliberate scope discipline."},
			{Label: "How it came together", Body: "I worked across the stack, tying the ML recommendation output into an Express/Prisma backend and a React surface, keeping each person's slice thin and reviewable."},
			{Label: "What I can stand behind", Body: "The idea landed in the Top 15 Best Capstone Projects. We proved a small, distributed team can ship a coherent AI product under pressure."},
		},
	},
	{
		Slug:  "spacelab",
		Title: "Spacelab",
		Year:  "2025",
		Tags:  []string{"Laravel", "PHP", "JavaScript"},
		Hero:  "Keeping school schedules, rooms, and teachers conflict-free without spreadsheet chaos.",
		Sections: []CaseStudySection{
			{Label: "Where it started", Body: "Building the school timetable on spreadsheets produced duplicates: two classes in one room, one teacher in two places. Spacelab was built to make conflicts structurally impossible."},
			{Label: "What boxed the work in", Body: "The scheduler is not a software person, so the tool had to enforce rules automatically without demanding technical effort. It also had to be fast to update."},
			{Label: "How it came together", Body: "Laravel modelled rooms, teachers, and class periods, and the JavaScript UI let staff drag and assign while the system flagged clashes as they happened."},
			{Label: "What I can stand behind", Body: "When every assignment is validated against the same rules, conflicts stop being discovered the morning of the class. The schedule finally stays honest."},
		},
	},
	{
		Slug:  "renshuu",
		Title: "Renshuu web",
		Year:  "2025",
		Tags:  []string{"React", "Laravel"},
		Hero:  "A job-search web app built during PKL for SMKN 1 Karawang students.",
		Sections: []CaseStudySection{
			{Label: "Where it started", Body: "Students finished internships but had no central place to find placement opportunities or track the application process. Renshuu (a play on the Indonesian word for training) gave them one."},
			{Label: "What boxed the work in", Body: "Built during PKL at CV Smartplus, the project had a firm timeline and a real student-facing audience — it had to feel genuinely usable, not a class exercise."},
			{Label: "How it came together", Body: "My team and I split a Laravel API and a React client. I focused on the search and application flow so a student could find a role and follow it to the end."},
			{Label: "What I can stand behind", Body: "It was shipped to the school's real context, used by real students. Working on a product with an actual deadline taught me the discipline of thin slices and honest shipping."},
		},
	},
}
