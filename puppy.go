package puppy

import "fmt"

func Bark() string {
	return "Bark"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

/*
firt add, commit and push all your changes
git tag v1.0.0 - attaches a tag to the commit
git push origin --tag - will push all tags that have been created
*/
func Versioning() string {
	fmt.Println("On version 3.0.0")
}
