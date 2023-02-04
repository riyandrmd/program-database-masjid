package models

type JadwalSholat struct {
	Id_jadwal int64  `gorm:"primaryKey;autoIncreament" json:"id_jadwal"`
	Hari      string `json:"hari"`
	Tanggal   string `json:"tanggal"`
	Shubuh    string `json:"shubuh"`
	Imsyak    string `json:"imsyak"`
	Zuhur     string `json:"zuhur"`
	Ashar     string `json:"ashar"`
	Maghrib   string `json:"maghrib"`
	Isya      string `json:"isya"`
}

type Inventaris struct {
	Id_Barang   int64  `gorm:"primaryKey;autoIncreament;" json:"id_barang"`
	Nama_Barang string `json:"nama_barang"`
	Jumlah      int64  `json:"jumlah"`
	Id_Ktg      int64  `json:"id_ktg"`
}

type Ktg_Inventaris struct {
	Id_Ktg     int64        `gorm:"primaryKey;autoIncreament;" json:"id_ktg"`
	Nama_Ktg   string       `json:"nama_ktg"`
	Inventaris []Inventaris `gorm:"foreignKey:Id_Ktg" constraint:"OnUpdate:CASCADE,OnDelete:SET NULL;" json:"inventaris"`
}
