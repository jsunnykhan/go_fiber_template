package redis

import "log"

func PublishMessage(r *redisService, channel string, message string) error {
	return r.client.Publish(ctx, channel, message).Err()
}

func SubscribeToEvents(r *redisService) {
	sub := r.client.Subscribe(ctx, "events")
	defer sub.Close()

	for {
		msg, err := sub.ReceiveMessage(ctx)
		if err != nil {
			log.Println("Error receiving message:", err)
			continue
		}
		log.Printf("Received message: %s", msg.Payload)
	}
}
