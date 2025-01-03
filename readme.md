<p align="center">
  <img src="animalized.png" />
</p>

# Introduction

This is open source realtime multiplayer game server project in Go. You can play with [animalized-client](https://github.com/BJS-kr/animalized-client). Initial purpose of this project is make a Go game server template. I couldn't find any descent, open sourced game server project in Go. So, I made my own(maybe this project bit gone too far for a template).

Unfortunately, I have no game server experience at production level. So if something looks unsuitable, please, let me know.

# Key Aspects
1. This is realtime simulation game
2. Avoiding potential blocks by [lock-free queue](queue/queue.go)(Yes, it could be dangerous. I'll explain it in Details)
3. TCP(only) server - This project uses [protobuf](https://protobuf.dev/)
4. [Actor](https://en.wikipedia.org/wiki/Actor_model), [Fan-out](https://en.wikipedia.org/wiki/Fan-out_(software)) pattern used for simple processing
5. [deterministic lockstep](https://www.linkedin.com/pulse/deterministic-lockstep-networking-demystified-zack-sinisi-jqrue) synchronization

# Game Rules(if you play with  [animalized-client](https://github.com/BJS-kr/animalized-client))

1. You can fire a fireball with key "a"
2. You can move characters using arrow keys
3. Fireball can hit a character or a wall
4. Wall will be weakened as fireball hit the wall
5. If wall state reached "vulnerable", fireball can destroy the wall
6. If a user hit player with fireball 10 times, game will be end.

# Flow
<p align="center">
  <img width=700 src="flow.jpg" />
</p>

# Details
### 1. How packet parsed?
1. Receive byte length
2. read amount of received byte length
3. parse into message
4. repeat

You can see the code at [packet_store.go](packet/packet_store.go).

### 2. How messages propagated?
All users are placed in actors. actors may vary, like lobby, room, game.
Actor sequentially receives all messages from owned users.
When actor receives message, it process the message, and fan-out to users.
actor may add message content, create new message, or drop the user message that don't have to be propagated.
<p align="center">
  <img width=900 src="propagate.jpg" />
</p>


### 3. Why Dispatcher do not directly send message but insert into queue?
dispatcher must guarantee same distribution performance. if distribution & sending occurs in single flow, performance cannot be guaranteed because user connection unstable or slow, overall distribution performance would not consistent.
so, distributor and sender should be separated and communicate via outgoing queue. This approach is based on common programming practice: separate pure functions and I/O operations. 

### 4. How user moved to another actor?
Users pass message to target channel. Message channel can be changed.
If user moves to another actor(e.g lobby -> room), message channel also changes.

### 5. How tick used?
This project uses ticks not only for in-game, but for all situation.
For example, lobby's tick rate is 200ms. Users in lobby receive queued messages every 200ms and consume.
Game's tick rate is 2ms, obviously for fast sync.

### 6. Why Lock-free Queue? (WHY NOT CHANNEL OR MUTEX?)
#### 6-a. Lock-free DS are not always fast and have pitfalls
Let's talk about general problems.

1. If spin lock used(I did), it can occur greater latency and excessive CPU consumption.
2. Simple is the best. Fancy algorithm tends to buggier and hard to change.
3. CAS implies ABA problem

I totally agree with those problems, but here's the reason.

**Answer to 1**. Lock-free is not suitable for high contention. It'll cause random latency and complex problems. In this project, lock-free queue is used for data receiving between dispatcher and sender only. It means there is no high contention but there are one queueing goroutine, and one dequeueing goroutine. Performance is quite expectable and beat up mutex case in performance always. You can check out the benchmark [here](queue/queue_test.go).

**Answer to 2**. I agree with that. Sincerely.

**Answer to 3**. Implementation and usage are monotonic in this project. ABA problem does not occurs in monotonic operation.

Not an answer for problems but another reason to use: It does not occur blocks. Let me explain it later(in "When it's better than channels").

#### 6-b. Why not sync.Map
It's simple. It does not guarantees order.
#### 6-c. When it's better than channels
I prefer channels in most cases. I have only one standard. If "block" makes sense, I use channels. 

A situation block makes sense is "expected" or "natural". For example, I used channel between Actor & Dispatcher. It makes sense because if Actor throughput itself is slow, Dispatcher receives message pace with actor. Users might have bad experience but it is a server problem and totally fair for players because all users receives message at a same pace.

On the contrary, Let's assume that some users have bad network condition. Message consume slowed down because writing to TCP connection is slowed down. Eventually, channel will be overflowed and block the whole operation. It is not "natural" because most of users were in good network condition.

Some might say channel can be buffered. Let's talk about that case. How much space would suitable for buffered channel, especially for fast stacking messages? How much should I buffer? I want to avoid channel blocking absolutely.

Some other might say channels can be length measured, and can prevent blocks. Here's the second problem. If you want to prevent block by measuring channel buffer length, you have to measure it every single iteration. Length of Go data structure can be cached ONLY if that ds used in local. If data referred outside, length cannot be cached and Go actually counts it.

#### 6-d. When it's better than mutex + slice

ADD: I implemented lock-free queue before move to Lock-step. This project currently uses ticks. So maybe mutex and slice would have better performance and simplicity. Even so, I did not changed because it works fine now and still can win against non-tick based operations

In my experience, using lock all around was a pure evil. I don't want to say it is a reason to prefer lock-free over mutex. It is just a SKILL ISSUE of mine. But still, lock-free can free me from deadlock hell a bit.

My main concern was performance. See 6-a.





