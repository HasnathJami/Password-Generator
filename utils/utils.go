package utils

import (
	"log"
)


func GetPasswordByCaesarCipher(password string) rune {

	var newPassword rune
	for index, _ := range password {
		newPassword += rune(password[index])
	}
	return newPassword * newPassword
}


func CheckSimpleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// func CheckRedisError(err *redis.StatusCmd) {
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func GeneratePassword(){
// 	 err := rdb.Set(ctx, "password", password, 0).Err()
// 	 utils.CheckSimpleError()
// 	 utils.CheckSimpleError(err)
// 	 //err = rdb.Set(ctx, "generated_password", utils.GetPasswordByCaesarCipher(password), 24*3600*time.Second).Err()
// 	 //utils.CheckSimpleError(err)
// }