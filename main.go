package main

import "fmt"

type Profile struct {
	Name   string
	Health int
	Power  int
	Exp    int
}

func MakeProfile(name string) Profile {
	switch name {
	case "Sasuke":
		return Profile{
			Name:   "Sasuke",
			Health: 200,
			Power:  100,
			Exp:    0,
		}
	case "Goku":
		return Profile{
			Name:   "Goku",
			Health: 400,
			Power:  300,
			Exp:    100,
		}
	case "Naruto":
		return Profile{
			Name:   "Naruto",
			Health: 150,
			Power:  200,
			Exp:    50,
		}
	default:
		return Profile{}
	}
}

func PowerUp(profile Profile, multiplier int) Profile {
	profile.Health = profile.Health + (profile.Health * multiplier)
	profile.Power = profile.Power + (profile.Power * multiplier)
	profile.Exp = profile.Exp + (profile.Exp * multiplier)
	return profile
}

func main() {
	ProfileNaruto := MakeProfile("Naruto")
	fmt.Println("Name:", ProfileNaruto.Name)
	fmt.Println("Health:", ProfileNaruto.Health)
	fmt.Println("Power:", ProfileNaruto.Power)
	fmt.Println("Exp:", ProfileNaruto.Exp)

	fmt.Println("===HEROS DIBUFF POWER UP===")

	powerUpNaruto := PowerUp(ProfileNaruto, 2)
	fmt.Println("Name:", powerUpNaruto.Name)
	fmt.Println("Health:", powerUpNaruto.Health)
	fmt.Println("Power:", powerUpNaruto.Power)
	fmt.Println("Exp:", powerUpNaruto.Exp)

	ProfileSasuke := MakeProfile("Sasuke")
	fmt.Println("Name:", ProfileSasuke.Name)
	fmt.Println("Health:", ProfileSasuke.Health)
	fmt.Println("Power:", ProfileSasuke.Power)
	fmt.Println("Exp:", ProfileSasuke.Exp)

	fmt.Println("===HEROS DIBUFF POWER UP===")

	powerUpSasuke := PowerUp(ProfileSasuke, 1)
	fmt.Println("Name:", powerUpSasuke.Name)
	fmt.Println("Health:", powerUpSasuke.Health)
	fmt.Println("Power:", powerUpSasuke.Power)
	fmt.Println("Exp:", powerUpSasuke.Exp)

	ProfileGoku := MakeProfile("Goku")
	fmt.Println("Name:", ProfileGoku.Name)
	fmt.Println("Health:", ProfileGoku.Health)
	fmt.Println("Power:", ProfileGoku.Power)
	fmt.Println("Exp:", ProfileGoku.Exp)

	fmt.Println("===HEROS DIBUFF POWER UP===")

	powerUpGoku := PowerUp(ProfileGoku, 1)
	fmt.Println("Name:", powerUpGoku.Name)
	fmt.Println("Health:", powerUpGoku.Health)
	fmt.Println("Power:", powerUpGoku.Power)
	fmt.Println("Exp:", powerUpGoku.Exp)
}
