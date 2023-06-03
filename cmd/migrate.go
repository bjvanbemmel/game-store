package cmd

type Migrate struct{}

func (m Migrate) Execute() error {
	if err := db.Connect(); err != nil {
		return err
	}

	return nil
}
