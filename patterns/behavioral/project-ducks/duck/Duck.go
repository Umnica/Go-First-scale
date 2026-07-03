package duck

type Duck struct {
    Name string  
}

func (d *Duck) SayHi() string {  
    return "Привет! Я утка " + d.Name
}