package main

func main() {
	works := Works{}
	storage := NewStorage[Works]("works.json")
	storage.Load(&works)

	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&works)

	storage.Save(works)
}
