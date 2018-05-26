package oneclient

import "github.com/fernandoocampo/katas/onedb"

var sequence int

type futureDB chan *result

type result struct {
	err  error
	data record
}

type record struct {
	id    int
	value string
}

// Create saves a new record in the repository.
// If something goes wrong it returns an error.
func Create(record string) (int, error) {

	done := createAsync(sequence, record)
	resp := <-done

	if resp.err != nil {
		return -1, resp.err
	}
	sequence++
	return resp.data.id, nil
}

// FindByID searches in the repository a record that
// contains the given key. If something goes wrong
// it returns an error.
func FindByID(id int) (string, error) {
	future := make(futureDB, 1)
	go func() {
		value, err := onedb.Find(id)
		future <- &result{
			err: err,
			data: record{
				value: value,
			},
		}
	}()
	resp := <-future

	if resp.err != nil {
		return "", resp.err
	}

	return resp.data.value, nil
}

func createAsync(id int, value string) futureDB {
	future := make(futureDB, 1)

	go func() {
		err := onedb.Save(id, value)
		future <- &result{
			err: err,
			data: record{
				id:    id,
				value: value,
			},
		}
	}()

	return future

}
