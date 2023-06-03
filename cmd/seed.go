package cmd

type Seed struct{}

func (s Seed) Execute() error {
	if err := db.Connect(); err != nil {
		return err
	}

	return nil
}
