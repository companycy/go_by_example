type Configuration struct {
	Val string
	Proxy
	
}

type Proxy struct {
	Address string
	Port    string
	
}

func main() {

	c := &Configuration{
		Val: "test",
		Proxy: Proxy{
			Address: "addr",
			Port:    "port",
			        
		},
		    
	}
	fmt.Println(c)
	
}
