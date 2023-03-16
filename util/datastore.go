package util

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"

	ordersv1 "github.com/froedevrolijk/grpc-invoicing/proto/orders/v1"
)

// LoadPbFromCsv loads the order data from csv into protobuf values
func LoadPbFromCsv(path string) ([]*ordersv1.Order, error) {
	items := make([]*ordersv1.Order, 0)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}

		var num int32
		if i, err := strconv.Atoi(row[1]); err == nil {
			num = int32(i)
		}

		c := &ordersv1.Order{
			Id:     row[0],
			Amount: num,
		}
		items = append(items, c)
	}
	return items, err
}
