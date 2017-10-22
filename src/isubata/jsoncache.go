package main

var jsonifyCache map[int64]User

func initJsonifyCache() error {
	jsonifyCache = make(map[int64]User, 2000)
	res := []User{}
	err := db.Get(&res, "SELECT id,name,display_name,avatar_icon FROM user")
	if err != nil {
		return err
	}

	for _, value := range res {
		jsonifyCache[value.ID] = value
	}

	return nil
}

func addJsonifyCache(id int64, name, display_name, avatar_icon string) {
	u := User{Name: name, DisplayName: display_name, AvatarIcon: avatar_icon}
	jsonifyCache[id] = u
}

func postJsonifyCache(id int64, display_name, avatar_icon string) {
	u := User{Name: jsonifyCache[id].Name, DisplayName: display_name, AvatarIcon: avatar_icon}
	jsonifyCache[id] = u
}
