namespace * call

struct Message {
  1:  i64 id,
  2:  string text
}

service CallService {
  Message call(1: Message m)
}