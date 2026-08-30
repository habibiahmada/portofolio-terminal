package sanitize

import "github.com/habibiahmada/habibiahmada-terminal/internal/data"

func Profile(p data.Profile) data.Profile {
	p.Name = Plain(p.Name)
	p.BoostName = Plain(p.BoostName)
	p.Title = Plain(p.Title)
	p.Location = Plain(p.Location)
	p.Email = Plain(p.Email)
	p.Availability = Plain(p.Availability)
	p.Website = Plain(p.Website)
	p.Employer = Plain(p.Employer)
	p.School = Plain(p.School)
	for i := range p.Stats {
		p.Stats[i].Value = Plain(p.Stats[i].Value)
		p.Stats[i].Label = Plain(p.Stats[i].Label)
	}
	return p
}

func Projects(ps []data.Project) []data.Project {
	out := make([]data.Project, len(ps))
	for i, p := range ps {
		out[i] = Project(p)
	}
	return out
}

func Project(p data.Project) data.Project {
	p.Name = Plain(p.Name)
	p.Slug = Plain(p.Slug)
	p.Year = Plain(p.Year)
	p.Description = Plain(p.Description)
	p.DescriptionID = Plain(p.DescriptionID)
	p.Tags = Strings(p.Tags)
	p.Stack = Strings(p.Stack)
	p.GitHub = Plain(p.GitHub)
	p.Live = Plain(p.Live)
	return p
}

func Skills(ss []data.Skill) []data.Skill {
	out := make([]data.Skill, len(ss))
	for i, s := range ss {
		out[i].Name = Plain(s.Name)
	}
	return out
}

func SkillCategories(cats []data.SkillCategory) []data.SkillCategory {
	out := make([]data.SkillCategory, len(cats))
	for i, c := range cats {
		out[i].Name = Plain(c.Name)
		out[i].Icon = Plain(c.Icon)
		out[i].Description = Plain(c.Description)
		out[i].Skills = Strings(c.Skills)
	}
	return out
}

func WorkExperience(ws []data.ExperienceWork) []data.ExperienceWork {
	out := make([]data.ExperienceWork, len(ws))
	for i, w := range ws {
		out[i].Period = Plain(w.Period)
		out[i].Role = Plain(w.Role)
		out[i].Company = Plain(w.Company)
		out[i].Location = Plain(w.Location)
		out[i].Badge = Plain(w.Badge)
		out[i].Details = Strings(w.Details)
	}
	return out
}

func Education(es []data.ExperienceEducation) []data.ExperienceEducation {
	out := make([]data.ExperienceEducation, len(es))
	for i, e := range es {
		out[i].Period = Plain(e.Period)
		out[i].Title = Plain(e.Title)
		out[i].School = Plain(e.School)
		out[i].Description = Plain(e.Description)
	}
	return out
}

func Certificates(cs []data.Certificate) []data.Certificate {
	out := make([]data.Certificate, len(cs))
	for i, c := range cs {
		out[i].Name = Plain(c.Name)
		out[i].Issuer = Plain(c.Issuer)
		out[i].Date = Plain(c.Date)
		out[i].URL = Plain(c.URL)
		out[i].Pinned = c.Pinned
	}
	return out
}

func Socials(ss []data.Social) []data.Social {
	out := make([]data.Social, len(ss))
	for i, s := range ss {
		out[i].Name = Plain(s.Name)
		out[i].URL = Plain(s.URL)
		out[i].Icon = Plain(s.Icon)
	}
	return out
}

func Companies(cs []data.Company) []data.Company {
	out := make([]data.Company, len(cs))
	for i, c := range cs {
		out[i].Name = Plain(c.Name)
	}
	return out
}

func Services(ss []data.Service) []data.Service {
	out := make([]data.Service, len(ss))
	for i, s := range ss {
		out[i].Number = Plain(s.Number)
		out[i].Category = Plain(s.Category)
		out[i].Title = Plain(s.Title)
		out[i].Description = Plain(s.Description)
	}
	return out
}

func ProcessSteps(ps []data.ProcessStep) []data.ProcessStep {
	out := make([]data.ProcessStep, len(ps))
	for i, p := range ps {
		out[i].Number = Plain(p.Number)
		out[i].Title = Plain(p.Title)
		out[i].Description = Plain(p.Description)
	}
	return out
}

func PressItems(ps []data.Press) []data.Press {
	out := make([]data.Press, len(ps))
	for i, p := range ps {
		out[i].Title = Plain(p.Title)
		out[i].Body = Plain(p.Body)
		out[i].CTALabel = Plain(p.CTALabel)
		out[i].URL = Plain(p.URL)
	}
	return out
}

func CaseStudy(cs data.CaseStudy) data.CaseStudy {
	cs.Slug = Plain(cs.Slug)
	cs.Title = Plain(cs.Title)
	cs.Year = Plain(cs.Year)
	cs.Tags = Strings(cs.Tags)
	cs.Hero = Plain(cs.Hero)
	cs.Live = Plain(cs.Live)
	cs.Source = Plain(cs.Source)
	for i := range cs.Sections {
		cs.Sections[i].Label = Plain(cs.Sections[i].Label)
		cs.Sections[i].Body = Plain(cs.Sections[i].Body)
	}
	return cs
}
