package seeders

import (
	"muslim_pro/config"
	"muslim_pro/models"
)

func SeederCaption() []models.Caption {

	captions := []models.Caption{
		{Email: "mirulid@gmail.com", Status: true, Baris1: "القدار", Baris2: "Allah plan is better than your dreams", Baris3: "يٰٓاَيُّهَا الَّذِيْنَ اٰمَنُوا اجْتَنِبُوْا كَثِيْرًا مِّنَ الظَّنِّۖ اِنَّ بَعْضَ الظَّنِّ اِثْمٌ وَّلَا تَجَسَّسُوْا وَلَا يَغْتَبْ بَّعْضُكُمْ بَعْضًاۗ اَيُحِبُّ اَحَدُكُمْ اَنْ يَّأْكُلَ لَحْمَ اَخِيْهِ مَيْتًا فَكَرِهْتُمُوْهُۗ وَاتَّقُوا اللّٰهَ ۗاِنَّ اللّٰهَ تَوَّابٌ رَّحِيْمٌ", Baris4: "Wahai orang-orang yang beriman! Jauhilah banyak dari prasangka, sesungguhnya sebagian prasangka itu dosa dan janganlah kamu mencari-cari kesalahan orang lain dan janganlah ada di antara kamu yang menggunjing sebagian yang lain. Apakah ada di antara kamu yang suka memakan daging saudaranya yang sudah mati? Tentu kamu merasa jijik. Dan bertakwalah kepada Allah, sesungguhnya Allah Maha Penerima tobat, Maha Penyayang."},
		{Email: "mirulid@gmail.com", Status: true, Baris1: "", Baris2: "Dua hal yang harus kita ingat :", Baris3: "Kebaikan orang lain terhadap kita dan Keburukan kita terhadap orang lain", Baris4: "by : Ustadz Handy Bonny"},
		{Email: "user4@example.com", Status: true, Baris1: "وَفِي السَّمَاءِ رِزْقُكُمْ وَمَا تُوعَدُونَ", Baris2: "And in the heaven is your provision and whatever you are promised", Baris3: "وَفِي السَّمَاءِ رِزْقُكُمْ وَمَا تُوعَدُونَ", Baris4: "Dan di langit terdapat rezekimu dan apa yang dijanjikan kepadamu."},
		{Email: "user5@example.com", Status: true, Baris1: "اِنَّ اللّٰهَ مَعَ الصَّابِرِيْنَ", Baris2: "Indeed, Allah is with the patient", Baris3: "يَا أَيُّهَا الَّذِينَ آمَنُوا اسْتَعِينُوا بِالصَّبْرِ وَالصَّلَاةِ ۚ إِنَّ اللَّهَ مَعَ الصَّابِرِينَ", Baris4: "Wahai orang-orang yang beriman! Mohonlah pertolongan (kepada Allah) dengan sabar dan salat. Sesungguhnya Allah beserta orang-orang yang sabar."},
		{Email: "user6@example.com", Status: true, Baris1: "وَمَا تَوْفِيقِي إِلَّا بِاللَّهِ", Baris2: "My success is only by Allah", Baris3: "وَمَا تَوْفِيقِي إِلَّا بِاللَّهِ ۚ عَلَيْهِ تَوَكَّلْتُ وَإِلَيْهِ أُنِيبُ", Baris4: "Dan tidak ada taufik bagiku melainkan dengan (pertolongan) Allah. Hanya kepada-Nya aku bertawakkal dan hanya kepada-Nya aku kembali."},
		{Email: "user7@example.com", Status: true, Baris1: "وَاصْبِرْ وَمَا صَبْرُكَ إِلَّا بِاللَّهِ", Baris2: "Be patient, and your patience is only through Allah", Baris3: "وَاصْبِرْ وَمَا صَبْرُكَ إِلَّا بِاللَّهِ ۚ وَلَا تَحْزَنْ عَلَيْهِمْ وَلَا تَكُ فِي ضَيْقٍ مِمَّا يَمْكُرُونَ", Baris4: "Bersabarlah, dan kesabaranmu itu hanya dengan pertolongan Allah, dan janganlah engkau bersedih hati terhadap (kekafiran) mereka dan jangan (pula) merasa sempit terhadap apa yang mereka tipu dayakan."},
	}

	for i := range captions {
		config.DB.Create(&captions[i])
	}

	return captions
}
