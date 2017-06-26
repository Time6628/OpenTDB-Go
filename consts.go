package OpenTDB_Go


const (
	general_knowledge = 9
	ent_books = 10
	ent_film = 11
	ent_music = 12
	ent_theatre = 13
	ent_television = 14
	ent_video_games = 15
	ent_board_games = 16
	ent_comics = 29
	ent_japanese = 31
	ent_cartoons_animation = 32
	sci_nature = 17
	sci_computers = 18
	sci_mathematics = 19
	sci_gadgets = 30
	mythology = 20
	sports = 21
	geography = 22
	history = 23
	politics = 24
	art = 25
	celebritites = 26
	animals = 27
	vehicles = 28
)

const (
	easy = "easy"
	medium = "medium"
	hard = "hard"
)

const (
	multiple = "multiple"
	boolean = "boolean"
)

const (
	base = "https://opentdb.com/api.php?amount="
	req_token = "https://opentdb.com/api_token.php?command=request"
	with_token = "&token="
	with_cat = "&category="
	with_difficulty = "&difficulty="
)