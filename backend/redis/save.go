package redis

import "context"

func Set(roomNum string, data []byte) error {
	if err := Cache.Set(context.TODO(), roomNum, data, 0).Err(); err != nil {
		return err
	}

	return nil
}

func Get(roomNum string) ([]byte, error) {
	data, err := Cache.Get(context.TODO(), roomNum).Result()
	if err != nil {
		return nil, err
	}
	return []byte(data), nil
}
