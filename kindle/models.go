package kindle

type Tabler interface {
	TableName() string
}

type BookInfo struct {
	Id      string `gorm:"primaryKey;column:id"`
	Asin    string `gorm:"column:asin"`
	Guid    string `gorm:"column:guid"`
	Lang    string `gorm:"column:lang"`
	Title   string `gorm:"column:title"`
	Authors string `gorm:"column:authors"`
}

func (BookInfo) TableName() string {
	return "BOOK_INFO"
}

type Word struct {
	Id        string `gorm:"primaryKey;column:id"`
	Word      string `gorm:"column:word"`
	Stem      string `gorm:"column:stem"`
	Lang      string `gorm:"column:lang"`
	Category  int    `gorm:"column:category"`
	Timestamp int64  `gorm:"column:timestamp"`
	ProfileId string `gorm:"column:profileid"`
}

type Lookup struct {
	Id        string `gorm:"primaryKey;column:id"`
	WordKey   string `gorm:"column:word_key;index"`
	BookKey   string `gorm:"column:book_key;index"`
	DictKey   string `gorm:"column:dict_key"`
	Pos       int    `gorm:"column:pos"`
	Usage     string `gorm:"column:usage"`
	Timestamp int64  `gorm:"column:timestamp"`

	Word     Word     `gorm:"foreignKey:Id;references:word_key"`
	BookInfo BookInfo `gorm:"foreignKey:Id;references:book_key"`
}

func (Lookup) TableName() string {
	return "LOOKUPS"
}

// TableName overrides the table name used by Word to `WORDS`
func (Word) TableName() string {
	return "WORDS"
}
