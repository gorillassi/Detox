package core

type Keypair struct {
	PublicKey  []byte `json:"pub"`
	PrivateKey []byte `json:"priv"`
}

type Post struct {
	AuthorID  string `json:"author_id"`
	Content   string `json:"content"`
	Timestamp int64  `json:"timestamp"`
	Sig       []byte `json:"sig"`
}