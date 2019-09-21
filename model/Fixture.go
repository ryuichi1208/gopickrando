func main() {
  log.SetFlags(log.Lshortfile)
  tempDir, err := ioutil.TempDir("", "golang-sample-echo-server.")
  pid := strconv.Itoa(os.Getpid())
  socket := tempDir + "/server." + pid
  listener, err := net.Listen("unix", socket)
  if err != nil {
    log.Printf("error: %v\n", err)
    return
  }
  if err := os.Chmod(socket, 0700); err != nil {
    log.Printf("error: %v\n", err)
    return
  }
  close := make(chan int)
  shutdown(listener, tempDir, close)
  fmt.Printf("GOLANG_SAMPLE_SOCK=%v;export GOLANG_SAMPLE_SOCK;\n", socket)
  fmt.Printf("GOLANG_SAMPLE_PID=%v;export GOLANG_SAMPLE_PID;\n", pid)
  server(listener)
  _ = <-close
}

func process(fd net.Conn) {
  defer fd.Close()
  for {
    buf := make([]byte, 512)
    nr, err := fd.Read(buf)
    if err != nil {
      break
    }
    data := buf[0:nr]
    fmt.Printf("Recieved: %v", string(data));
    _, err = fd.Write(data)
    if err != nil {
      log.Printf("error: %v\n", err)
      break
    }
  }
}
