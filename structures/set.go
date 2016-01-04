package structures


type Set struct{
	
	Set map[string] bool

}

func(set *Set) Add(i string)  bool {

	_, found := set.Set[i]

	set.Set[i] = true

	return !found
}
