package domain

type User struct {
	ID       string `firestore:"id"`
	Name     string `firestore:"name" binding:"required"`
	Email    string `firestore:"email" binding:"required,email"`
	Password string `firestore:"password" binding:"required"`
}

type Login struct{
	Email    string `firestore:"email"`
	Password string `firestore:"password"`
}
