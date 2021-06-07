package formatter

func ValidateAndFormatFrontMatterVersion(frontMatterVersion string) (string, error) {
	// TODO:
	if err := validateFrontMatterVersion(frontMatterVersion); err != nil {
		return "", err
	}
	return formatFrontMatterVersion(frontMatterVersion), nil
}

func validateFrontMatterVersion(frontMatterVersion string) error {
	// TODO:
	return nil
}

func formatFrontMatterVersion(validatedFrontMatterVersion string) string { // NOTE: the arg must be validated.
	// TODO:
	return ""
}

func ValidateAndFormatTitle(title string) (string, error) {
	// TODO:
	if err := validateTitle(title); err != nil {
		return "", err
	}
	return formatTitle(title), nil
}

func validateTitle(title string) error {
	// TODO:
	return nil
}

func formatTitle(validatedTitle string) string { // NOTE: the arg must be validated.
	// TODO:
	return ""
}

func ValidateAndFormatDrafted(drafted string) (string, error) {
	// TODO:
	if err := validateDrafted(drafted); err != nil {
		return "", err
	}
	return formatDrafted(drafted), nil
}

func validateDrafted(drafted string) error {
	// TODO:
	return nil
}

func formatDrafted(validatedDrafted string) string { // NOTE: the arg must be validated.
	// TODO:
	return ""
}

func ValidateAndFormatCreated(created string) (string, error) {
	// TODO:
	if err := validateCreated(created); err != nil {
		return "", err
	}
	return formatCreated(created), nil
}

func validateCreated(created string) error {
	// TODO:
	return nil
}

func formatCreated(validatedCreated string) string { // NOTE: the arg must be validated.
	// TODO:
	return ""
}

func ValidateAndFormatLastUpdated(lastUpdated string) (string, error) {
	// TODO:
	if err := validateLastUpdated(lastUpdated); err != nil {
		return "", err
	}
	return formatLastUpdated(lastUpdated), nil
}

func validateLastUpdated(lastUpdated string) error {
	// TODO:
	return nil
}

func formatLastUpdated(validatedLastUpdated string) string { // NOTE: the arg must be validated.
	// TODO:
	return ""
}

func ValidateAndFormatLastChecked(lastChecked string) (string, error) {
	// TODO:
	if err := validateLastChecked(lastChecked); err != nil {
		return "", err
	}
	return formatLastChecked(lastChecked), nil
}

func validateLastChecked(lastChecked string) error {
	// TODO:
	return nil
}

func formatLastChecked(validatedLastChecked string) string { // NOTE: the arg must be validated.
	// TODO:
	return ""
}

func ValidateAndFormatTags(tags []string) ([]string, error) {
	// TODO:
	if err := validateTags(tags); err != nil {
		return nil, err
	}
	return formatTags(tags), nil
}

func validateTags(tags []string) error {
	// TODO:
	return nil
}

func formatTags(validatedTags []string) []string { // NOTE: the arg must be validated.
	// TODO:
	return nil
}

func ValidateAndFormatID(id string) (string, error) {
	// TODO:
	if err := validateID(id); err != nil {
		return "", err
	}
	return formatID(id), nil
}

func validateID(id string) error {
	// TODO:
	return nil
}

func formatID(validatedID string) string { // NOTE: the arg must be validated.
	// TODO:
	return ""
}
