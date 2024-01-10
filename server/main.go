package main

func main() {
	//l1, err := reuse.Listen("tcp", "0.0.0.0:8888")
	//if err != nil {
	//	panic(err)
	//}
	//bt := make([]byte, 1024)
	//for {
	//	conn, err := l1.Accept()
	//	if err != nil {
	//		panic(err)
	//	}
	//	go func() {
	//		for {
	//			n, err := conn.Read(bt)
	//			if err != nil {
	//				panic(err)
	//			}
	//			body := bt[:n]
	//			fmt.Println("read", body, "from", conn.RemoteAddr())
	//
	//		}
	//	}()
	//}
	const (
		a = iota
		b = 5
		c
		d = iota
	)
}
