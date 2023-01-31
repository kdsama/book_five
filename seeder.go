package main

import "github.com/kdsama/book_five/service"

func categorySeeder(categoryservice *service.CategoryService) {
	categoryservice.SaveCategory("Adventure stories", []string{})
	categoryservice.SaveCategory("Classics", []string{})
	categoryservice.SaveCategory("Crime", []string{})
	categoryservice.SaveCategory("Fairy tales, fables, and folk tales", []string{})
	categoryservice.SaveCategory("Fantasy", []string{})
	categoryservice.SaveCategory("Historical fiction", []string{})
	categoryservice.SaveCategory("Horror", []string{})
	categoryservice.SaveCategory("Humour and satire", []string{})
	categoryservice.SaveCategory("Literary fiction", []string{})
	categoryservice.SaveCategory("Mystery", []string{})
	categoryservice.SaveCategory("Poetry", []string{})
	categoryservice.SaveCategory("Plays", []string{})
	categoryservice.SaveCategory("Romance", []string{})
	categoryservice.SaveCategory("Science fiction", []string{})
	categoryservice.SaveCategory("Short stories", []string{})
	categoryservice.SaveCategory("Thrillers", []string{})
	categoryservice.SaveCategory("War", []string{})
	categoryservice.SaveCategory("Women’s fiction", []string{})
	categoryservice.SaveCategory("Young adult", []string{})
	categoryservice.SaveCategory("Fiction", []string{"Adventure stories", "Classics", "Crime", "Fairy tales, fables, and folk tales", "Fantasy", "Historical fiction", "Horror", "Humour and satire", "Literary fiction", "Mystery", "Poetry", "Plays", "Romance", "Science fiction", "Short stories", "Thrillers", "War", "Women’s fiction", "Young adult"})
	categoryservice.SaveCategory("Autobiography and memoir", []string{})
	categoryservice.SaveCategory("Biography", []string{})
	categoryservice.SaveCategory("Essays", []string{})
	categoryservice.SaveCategory("Self-help", []string{})
	categoryservice.SaveCategory("Non-fiction", []string{"Autobiography and memoir", "Biography", "Essays", "Non-fiction novel", "Self-help"})
}
