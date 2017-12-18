package logstd

import "log"

func init() {
	log.SetPrefix("luzq:")
	log.SetFlags(log.Ldate | log.Llongfile | log.Lmicroseconds)
}

func Run() {
	log.Println("mmp")
	log.Fatalln("fmmp")
	log.Panicln("pmmp")

}
