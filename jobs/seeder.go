package jobs

import "github.com/kdsama/book_five/service"

func CategorySeeder(categoryservice *service.CategoryService) {
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
	categoryservice.SaveCategory("Computer Science", []string{})
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

func SaveBooks(bookservice service.BookServiceInterface) {

	type bookSeedStruct struct {
		name        string
		authors     []string
		co_authors  []string
		audiobooks  []string
		ebook_urls  []string
		hard_copies []string
		categories  []string
	}
	listing := []bookSeedStruct{
		{"Clean Code", []string{"Robert C.Martin"}, []string{}, []string{}, []string{}, []string{}, []string{"Computer Science"}},
		{"Chasing Excellence: A Story About Building the World's Fittest Athletes", []string{"Ben Bergeron"}, []string{}, []string{}, []string{}, []string{}, []string{"Self-help"}},
		{"1984", []string{"George Orwell"}, []string{}, []string{}, []string{}, []string{}, []string{"Fiction"}},
		{"Never Finished", []string{"David Goggins"}, []string{}, []string{}, []string{}, []string{}, []string{"Self-help"}},
		{"Finding Ultra", []string{"Rich Roll"}, []string{}, []string{}, []string{}, []string{}, []string{"Self-help"}},
		{"Astrophysics for Young People in a Hurry", []string{"Gregory Mone"}, []string{"Gabrielle de Cuir", "Neil deGrasse Tyson"}, []string{}, []string{}, []string{}, []string{"Non-fiction"}},
		{"Can't Hurt Me", []string{"David Goggins"}, []string{}, []string{}, []string{}, []string{}, []string{"Self-help"}},
		{"Art of Seduction", []string{"Robert Green"}, []string{}, []string{}, []string{}, []string{}, []string{"Self-help"}},
		{"12 Rules of Life", []string{"Jordan B.Peterson"}, []string{}, []string{}, []string{}, []string{}, []string{"Self-help"}},
	}
	for i := range listing {
		bookservice.SaveBook(listing[i].name, "", listing[i].authors, listing[i].co_authors, listing[i].audiobooks, listing[i].ebook_urls, listing[i].hard_copies, listing[i].categories)
	}
}
