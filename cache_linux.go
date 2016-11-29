package main

func flushPages() {
	//Shell("free")
	Shell("sudo", "sh", "-c", "echo 3 > /proc/sys/vm/drop_caches")
	//Shell("free")

}
