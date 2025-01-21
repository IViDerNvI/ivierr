package ivierr

import "reflect"

func Is(err, target error) bool {
	return !reflect.DeepEqual(kmpSearch(err.Error(), target.Error()), []int{})
}
