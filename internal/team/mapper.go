package team

// ToTeam -
func ToTeam(teamDTO DTO) Team {
	return Team{
		ID:        teamDTO.ID,
		Name:      teamDTO.Name,
		Status:    teamDTO.Status,
		CreatedAt: teamDTO.CreatedAt,
		UpdatedAt: teamDTO.UpdatedAt,
	}
}

// ToTeamDTO -
func ToTeamDTO(team Team) DTO {
	return DTO{ID: team.ID, Name: team.Name, Status: team.Status, CreatedAt: team.CreatedAt, UpdatedAt: team.UpdatedAt}
}

// ToTeamDTOs -
func ToTeamDTOs(teams []Team) []DTO {
	teamdtos := make([]DTO, len(teams))

	for i, itm := range teams {
		teamdtos[i] = ToTeamDTO(itm)
	}

	return teamdtos
}
