package main

import (
	"image/color"
	"math/rand"
	"strconv" //string convertion
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {

	var races = [50]string{
		"Aarakocra", "Aasimar",
		"Centaure", "Changelin",
		"Dhampire", "Demi-elf", "Demi-orc", "Drakéide (dragonborn)",
		"Elf",
		"Fée", "Firbolg",
		"Genasi", "Giff", "Gith", "Gnome", "Gobelours (bugbear)", "Goblin", "Goliate", "Guerre-Forgée", "Grung",
		"Hadozee", "Hilin (owlin)", "Hobgoblin", "Hobbit", "Humain",
		"Kalashtar", "Kenku", "Kobold", "Kor",
		"Lézardin (lizardfolk)", "Leonin", "Lièvron (Harengon)", "Loxodon",
		"Minotaure",
		"Naga", "Nain",
		"Orc",
		"Plasemoïde",
		"Sang-Maudit (hexblood)", "Sirène (merfolk)", "Satyr", "Shader-Kai",
		"Tabaxi", "Thri-Kreen", "Tiefling", "Tortle", "Triton",
		"Vedalken", "Verdan",
		"Yuan-Ti"}

	var classes = [13]string{
		"Artificier",
		"Barbare",
		"Bard",
		"Cleric",
		"Combattant",
		"Démoniste",
		"Druid",
		"Forestier",
		"Magicien",
		"Moine",
		"Paladin",
		"Roublard",
		"Sorcier"}

	var origines = [32]string{
		"Acolyte", "Anthropologist", "Archéologiste", "Artisan de guilde", "Athlète",
		"Charlatan", "Chasseur de prime", "Chevalier", "Contrebandier", "Courtier", "Criminèle",
		"Divértisseur",
		"Espion",
		"Gladiateur",
		"Hantée", "Hermit", "Héritier", "Héro du peuple",
		"Inspecteur",
		"Lointain voyageur",
		"Marchand", "Marchand ratée", "Marin", "Mercenaire", "Misérable",
		"Noble",
		"Parieur", "Perdu du Fey", "Pécheur", "Pirate",
		"Sage", "Soldat"}

	rand.Seed(time.Now().UnixNano()) //declare the seed used to make the random number

	application := app.New()
	icon, _ := fyne.LoadResourceFromPath("./d20changed.png")
	application.SetIcon(icon)

	window := application.NewWindow("Generateur de perso")
	var content fyne.Container = create_layout(races[:], classes[:], origines[:])

	button := widget.NewButton("Relance", func() {
		content = create_layout(races[:], classes[:], origines[:])
		window.Content().Refresh()
	})

	content2 := container.New(layout.NewVBoxLayout(), button, &content)

	window.SetContent(content2)
	window.CenterOnScreen()
	window.ShowAndRun()

	// fmt.Println("Race :", race, "("+bonuses_race(race)+")")
	// fmt.Println("Classe :", classe, "("+sous_classe(classe)+")")
	// fmt.Println("Histoire :", histoire, "("+bonuses_histoire(histoire)+")")
	// fmt.Println("Stats :", create_stat(), create_stat(), create_stat(), create_stat(), create_stat(), create_stat())
}

//TODO
func create_layout(races []string, classes []string, origines []string) fyne.Container {

	// randomly chooses one race, class and background
	var race string = choose_randomly(races[:])
	var classe string = choose_randomly(classes[:])
	var origine string = choose_randomly(origines[:])

	//first i create the label of the box, then i create a canvas object that accepts wrapping and line breaks and finaly i put this object in a scroll box
	race_label := canvas.NewText(race, color.Black)
	race_label.TextSize = 16
	race_label.TextStyle.Bold = true
	race_content := widget.NewLabel(bonuses_race(race))
	race_content.Wrapping = fyne.TextWrapWord
	race_content_scroll := container.NewScroll(race_content)
	race_content_scroll.SetMinSize(fyne.NewSize(500, 150))

	class_label := canvas.NewText(classe, color.Black)
	class_label.TextSize = 16
	class_label.TextStyle.Bold = true
	class_content := widget.NewLabel(sous_classe(classe))
	class_content.Wrapping = fyne.TextWrapWord
	class_content_scroll := container.NewScroll(class_content)
	class_content_scroll.SetMinSize(fyne.NewSize(500, 150))

	origine_label := canvas.NewText(origine, color.Black)
	origine_label.TextSize = 16
	origine_label.TextStyle.Bold = true
	origine_content := widget.NewLabel(bonuses_origine(origine))
	origine_content.Wrapping = fyne.TextWrapWord
	origine_content_scroll := container.NewScroll(origine_content)
	origine_content_scroll.SetMinSize(fyne.NewSize(500, 150))

	stats_label := canvas.NewText("stats ", color.Black)
	stats_label.TextSize = 16
	stats_label.TextStyle.Bold = true
	stats_content := widget.NewLabel(roll_stat() + roll_stat() + roll_stat() + roll_stat() + roll_stat() + roll_stat())

	race_container := put_in_container(race_label, race_content_scroll)

	class_container := put_in_container(class_label, class_content_scroll)

	origine_container := put_in_container(origine_label, origine_content_scroll)

	stats_container := put_in_container(stats_label, stats_content)

	conteneur := container.New(layout.NewVBoxLayout(), &race_container, &class_container, &origine_container, &stats_container)

	return *conteneur
}

func put_in_container(label fyne.CanvasObject, content fyne.CanvasObject) fyne.Container {

	spacer := widget.NewLabel("   ")

	label_container := container.New(layout.NewHBoxLayout(), label, spacer)
	content_container := container.New(layout.NewHBoxLayout(), spacer, content)
	container := container.New(layout.NewVBoxLayout(), label_container, content_container, spacer)

	return *container
}

//chooses a random element in an array
func choose_randomly(array []string) string {

	random_choice := array[rand.Intn(len(array))] //selects a random variable in the array
	return random_choice
}

//throws four 6 sided dice, gets rid of the lowest and sums up the rest (returns a string)
func roll_stat() string {

	var stat int = 0
	var smallest_die int = 7

	for i := 0; i < 4; i++ {

		die := rand.Intn(6) + 1

		if die < smallest_die {
			smallest_die = die
		}

		stat += die
	}
	stat -= smallest_die
	return_stat := strconv.Itoa(stat) + " " //strconv.itoa converts the int to a string
	return return_stat
}

//selects the bonus of the race
func bonuses_race(race string) string {
	var bonus string = ""

	switch race {
	case "Aarakocra":
		bonus = "Dex + 2, Sag + 1 \n\n25 pieds, 50 en vol \n\nvol (impossible si armure moyenne ou lourde est portée) \n\nattaque sans arme = 1d4 + force \n\nAarakocra et Auran"
	case "Aasimar":
		bonus = "Cha + 2, Sag + 1 \n\n30 pieds \n\nvision nocturne \n\nresistant au dégat necrotic et radiant \n\nconnait le sort \"lumière\" puis au niveau 3 \"restoration mineur\"\npuis au niveau 5\"lumière du jour\", peuvent être lancée 1 fois par jour (charisme) \n\nCommun et Celestial"
	case "Centaure":
		bonus = "For + 2, Sag + 1 \n\n40 pieds \n\nsi vous parcourez 30 pieds avant votre attaque vous pouvez utilisez votre action bonus pour faire une autre attaque \n\n attaque sans arme = 1d4 + force \n\ndouble la charge que vous pouvez porter mais difficultée à grimper \n\nchoisit 1: domptage, medecine, nature, survie \n\nCommun et Sylvain"
	case "Changelin":
		bonus = "Cha + 2, X + 1 \n\n30 pieds \n\npeut changer de forme mais pas au point de changer de catégorie de taille et les stats ne changent pas \n\nchoisit 2 : mensonge, lire le visage, intimidation, persuasion \n\nCommun et deux au choix "
	case "Dhampire":
		bonus = "X + 2, Y + 1 \n\n35 pieds, bon grimpeur \n\nvision nocturne \n\nà partir du niveau 3 peut se déplacer sur les murs comme si le sort \"grimpe d'araignée\" était lancée \n\nattaque sans arme = 1d4 + constitution, avantage sur les cibles qui sont sous la moitié des pv, au choix: regene les dégats infligés,ajoute les dégats infligé comme bonus au prochain dée lancé \n\nCommun et un autre au choix"
	case "Demi-elf":
		bonus = "Cha + 2, X + 1, Y + 1 \n\n30 pieds \n\nvision nocturne \n\navantage sur sauvetage contre charme magique et impossible d'être endormi par la magie \n\ndevient pro dans 2 talents  \n\nCommun, Elf, 1 autre"
	case "Demi-orc":
		bonus = "For + 2, Con + 1 \n\n30 pieds \n\nvision nocturne \n\nproeficient en intimidation \n\n1 fois par jour quand tombe à 0 pv tombe à 1 pv à la place \n\nquand inflige coup critique ajoute 1 dée en plus sur les dégats \n\nCommun et Orc"
	case "Drakéide (dragonborn)":
		bonus = "For + 2, Cha + 1 \n\n30 pieds \n\npeut cracher des flames (ou autre type de dégat par rapport à la couleur de peau) dans une colonne de 5 par 30 pieds ou un cone de 15 pieds, sauvetage = 8 + con + pro, 2d6 puis au niv6 3d6 puis au niv11 4d6 puis au niv16 5d6 \n\nrésitant au type de dégat du crachat \n\nCommun et Draconique"
	case "Elf":
		bonus = "Dex + 2 \n\n30 pieds \n\nvision nocturne \n\nproeficient en perception \n\npas besoin de dormir mais doit méditer pendant 4 heures \n\n Commun et Elf"
	case "Fée":
		bonus = "X + 2, Y + 1 \n\n30 pieds \n\n petit \n\napprend le sort \"truc de druide\", au niv 3 peut lancer \"feu de fée\", au niv 5 peut lancer \"elargir, raptisir\", ne peuvent être lancée qu'une fois par jour (charisme ou sagesse) \n\nvol (impossible si armure moyenne ou lourde est portée) \n\nCommun et 1 autre au choix"
	case "Firbolg":
		bonus = "Sag + 2, For + 1 \n\n30 pieds \n\npeut lancer \"detecter la magie\" et \"deguissement\", déguissement peut modifer la taille de 3 pieds \n\npeut devenir invisible jusqu'au début de son prochain tout un nombre de fois = à la proéficiense \n\ndouble le poid soulevable \n\nCommun, avec les animaux, 1 autre au choix"
	case "Genasi":
		bonus = "Con + 2 \n\n30 pieds"
	case "Giff":
		bonus = "X + 2, Y + 1 \n\n30 pieds, nage bien \n\nrelance les 1 sur les dées de dégats \n\navantage sur les jets de force, double le poids soulevable \n\nCommun et 1 autre au choix"
	case "Gith":
		bonus = "Int + 1 \n\n30 pieds"
	case "Gnome":
		bonus = "Int + 2 \n\n25 pieds \n\nvision nocturne \n\navantage sur les sauvetage contre la magie \n\nCommun et Gnome"
	case "Gobelours (bugbear)":
		bonus = "For + 2, Dex + 1 \n\n30 pieds \n\nvision nocturne \n\n+1 case de porté d'attaque au càc \n\ndouble le poid soulevable \n\npro en discrétion \n\n+2d6 dégats sur les attaques surprises \n\nCommun et Goblin"
	case "Goblin":
		bonus = "Dex + 2, Con + 1 \n\n30 pieds \n\npetit \n\nvision nocturne \n\n1 fois par jour peut ajouter son niveau à ses dégats si la cible est plus grande \n\npeut désengager et se cacher en action bonus \n\nCommun et Goblin"
	case "Goliate":
		bonus = "For + 2, Con + 1 \n\n30 pieds \n\npro en athlétisme \n\n1 fois par jour peut réduire les dégat subit de 1d12 + con \n\ndouble le poids soulevable \n\nrésitant au dégat de froids \n\nCommun et Géant"
	case "Guerre-Forgée":
		bonus = "Con + 2, X + 1 \n\n30 pieds \n\nrésistant au poison, pas besoin de manger,boire,dormire,respirer, immuniser au sommeil et au maladie \n\ndoit prendre des repos long mais n'est pas inconcient pendant ceux-ci \n\n+1 armure \n\n+1 pro talent, +1 pro outils \n\nCommun et un au choix "
	case "Grung":
		bonus = "Dex + 2, Con + 1 \n\n25 pieds, nage bien \n\npetit \n\nproéficient en mensonge \n\namphibien \n\nimmunisé au poison \n\npeau sécrete un poison DC12 2d4 dégat poison \n\nsaute 25pieds \n\ndoit s'immerger dans l'eau 1 heures par jour ou souffire 1 niveau de fatigue \n\nGrung"
	case "Hadozee":
		bonus = "X + 2, Y + 1 \n\n30 pieds, grimpe bien \n\npetit ou moyen \n\nplane (impossible en armure lourde) \n\nCommun et 1 au choix"
	case "Hilin (owlin)":
		bonus = "X + 2, Y + 1 \n\n30 pieds \n\nvision nocturne \n\nvol (impossible si armure moyenne ou lourde est portée) \n\npro en discrétion \n\nCommun et 1 au choix"
	case "Hobgoblin":
		bonus = "Con + 2, Int + 1 \n\n30 pieds \n\nvision nocturne \n\n+2 pro arme \n\nsi rate un jet peut ajouter un bonus = au nombre d'allié adjacent (max 5) \n\nCommun et Goblin"
	case "Hobbit":
		bonus = "Dex + 2 \n\n25 pieds \n\npetit \n\nrelance les 1 \n\navantage contre la peur \n\nCommun et Hobbite"
	case "Humain":
		bonus = "+ 1 toutes les stats \n\n30 pieds \n\nCommun et 1 au choix"
	case "Kalashtar":
		bonus = "Sag + 2, Cha + 1 \n\n30 pieds \n\navantage sur jet de sagesse \n\nrésitant au dégat psychic \n\ntélépatie sur cible visible \n\nimmuniser au sort qui font réver \n\nCommun, Quori et 1 au choix"
	case "Kenku":
		bonus = "Dex + 2, Sag + 1 \n\n30 pieds \n\npeut mimiquer les voix et les manurismes \n\nchoisit 2 : acrobatie, mensonge, discrétion, jeux de main \n\nCommun et Auran mais peut seulement parler via mimicerie"
	case "Kobold":
		bonus = "Dex + 2 \n\n30 pieds \n\npetit \n\n1 fois par jour peut donner avantage à ces allié en se mettant en boule et commencer à pleurer \n\navantage sur attaque si allié autour de la cible \n\ndéavantage à la vision quand sous le soleil \n\nCommun et Draconique"
	case "Kor":
		bonus = "Dex + 2, Sag + 1 \n\n30 pieds, grimpe bien \n\navantage contre la peur \n\nproéficient en acrobatie et athlétisme \n\nrelance les 1 \n\nCommun et language des signes"
	case "Lézardin (lizardfolk)":
		bonus = "Con + 2, Sag + 1 \n\n30 pieds, nage bien \n\nattaque sans arme = 1d6 + force \n\npeut transformer un cadavre en arme \n\npeut retenir sa réspiration jusqu'à 15 minutes \n\nchoisit 2 : domptage, nature, perception, discrétion, survie \n\nsans armure = 13 + dex \n\naction bonus peut mordre et gagner point de vie temporaire = au dégat + con \n\nCommun et Draconique"
	case "Leonin":
		bonus = "Con + 2, For + 1 \n\n35 pieds \n\nvision nocturne \n\nattaque sans arme = 1d4+force \n\nchoisit 1 : athlétisme, intimidation, perception, survie \n\n1 fois par jour action bonus terrifier ennemies 10 pieds autour de soi si rate sauvetage de sagesse = 8+pro+con \n\nCommun et Leonin"
	case "Lièvron (Harengon)":
		bonus = "X + 2, Y + 1 \n\n30 pieds \n\npetit ou moyen \n\n+pro sur initiative \n\npro perception \n\nen réaction peut ajouter 1d4 au jet de dex \n\nen action bonus peut sauter une distance = 5*pro en pieds \n\nCommun et 1 au choix"
	case "Loxodon":
		bonus = "Con + 2, Sag + 1 \n\n30 pieds \n\ndouble le poid soulevalbe \n\navantage contre charme et peur \n\nsans armure = 12 + con \n\ntrompe sert de 3ième main \n\navantage sur jet pour sentir(nez) \n\nCommun et Loxodon"
	case "Minotaure":
		bonus = "For + 2, Con + 1 \n\n30 pieds \n\nattaque sans arme = 1d6 + force \n\nsi dash peut attaque \n\nsi attaque touche peut tenter de pousser ennemie avec sau de force DC = 8+pro+for \n\nchoisit 1 : intimidation, persuasion \n\nCommun et Minotaure"
	case "Naga":
		bonus = "Con + 2, Int + 1 \n\n30 pieds \n\naction bonus peut augmenter vitesse de 5 pieds \n\nattaque sans arme = 1d4+for,sauv. con DC 8+pro+con si rate 1d4 poison ou\n1d6+for,sauv. for DC 8+pro+for si rate cible constrainte \n\nimmuniser au poison \n\npeut créer des poisons \n\nCommun et Naga"
	case "Nain":
		bonus = "Con + 2 \n\n25 pieds \n\nvision nocturne \n\navantage et réistance au poison \n\npro hache et marteau + choisit 1 : forge, brasserie, masson \n\ndouble bonus sur jet sur histoire de la pierre \n\nCommun et Nain"
	case "Orc":
		bonus = "For + 2, Con + 1 \n\n30 pieds \n\nvision nocturne \n\naction bonus peut se déplacer d'une distance = vitesse/2 \n\nchoisit 2 : domptage, lire le visage, intimidation, medecine, perception, survie \n\ndouble le poid soulevable \n\nCommun et Orc"
	case "Plasemoïde":
		bonus = "X + 2, Y + 1 \n\n30 pieds \n\npetit ou moyen \n\nsi rien d'équipé peut passer dans des trous de 1cm de diamètre \n\nvision nocturne \n\npeut retenir soufle pendant 1 heure \n\nrésistant au poison et acide \n\npeut détacher une partie de son corp pour effectuer des tâches(ne peut pas attaquer, utiliser d'object magique ou porter des charges plus grande que 5kg) \n\nCommun et 1 au choix"
	case "Sang-Maudit (hexblood)":
		bonus = "X + 2, Y + 1 \n\nbonus de la race maudite"
	case "Sirène (merfolk)":
		bonus = "Cha + 1 \n\n30 pieds nage bien \n\namphibien \n\nCommun et Sirène"
	case "Satyr":
		bonus = "X + 2, Y + 1 \n\n35 pieds \n\nattaque sans arme = 1d4+for \n\navantage sur jet contre magie \n\ndistance saut += 1d8 \n\npro performance et persuasion + 1 instrument \n\nCommun et Sylvain"
	case "Shader-Kai":
		bonus = "X + 2, Y + 1 \n\n30 pieds \n\nen action bonus peut se teleporter sur 30 pieds et devient résitant au dégat jusqu'au début de son prochain tour \n\nvision nocturne \n\navantage contre charme \n\npro perception \n\nrésistant au dégat nécrotique \n\npas besoin de dormire mais dois méditer 4 heures \n\nCommun et 1 au choix"
	case "Tabaxi":
		bonus = "Dex + 2, Cha + 1 \n\n 30 pieds, grimpe bien \n\nattaque sans arme = 1d6+for \n\npro acrobatie et discrétion \n\n1 fois par jour peut doubler vitesse pendant 1 tour \n\nCommun et 1 au choix"
	case "Thri-Kreen":
		bonus = "X + 2, Y + 1 \n\n30 pieds \n\npetit ou moyen \n\nsans armure = 13+dex, peut se camoufler \n\nvision nocturne \n\n+2 bras \n\npas besoin de dormir \n\ntelepatie jusqu'à 120 pieds \n\nCommun et 1 au choix"
	case "Tiefling":
		bonus = "Cha + 2, Int + 1 \n\n30 pieds \n\nvision nocturne \n\nrésistant au feu \n\napprend \"thaumaturgie\", a partir du niveau 3 peut lancer \"rétribution infernal\" (niv 2) 1 fois par jour, a partir du niveau 5 peut lancer\"tenebre\" 1 fois par jour, (utilise charisme) \n\nCommun et Infernal"
	case "Tortle":
		bonus = "For + 2, Sag + 1 \n\n30 pieds \n\nattaque sans arme = 1d4+for \n\npeut retenir son soufle pendant 1 heures \n\nsans armure = 17 \n\nen action peut rentrer dans sa carapace et gagner 4 armure \n\npro survie \n\nCommun et Aquan"
	case "Triton":
		bonus = "For + 1, Con + 1, Cha + 1 \n\n30 pieds, nage bien \n\namphibien \n\npeut lancer \"brouillard\" 1 fois par jour, à partir du niveau 3 peut lancer \"courant d'air\" 1 fois par jour, à partir du niveu 5 peut lancer \"mur d'eau\" 1 fois par jour (charisme) \n\nvision nocture \n\nrésistan au froid \n\nCommun, Primordiale et avec les poisssons"
	case "Vedalken":
		bonus = "Int + 2, Sag + 1 \n\n30 pieds \n\navantage sur jet d'intelligence, sagesse et charsime \n\nchoisit 1 : Arcane, histoire, investigation, medecine, performance, jeux de main et ajout 1d4 sur les jets de ce talent \n\npeut respirer sous l'eau pendant 1 heure \n\nCommun, Vedalkain et 1 au choix"
	case "Verdan":
		bonus = "Cha + 2, Con + 1 \n\n30 pieds \n\nrelance les 1 et 2 sur soin de Hit Die \n\ntélépatie sur 30 pieds \n\npro persuasion \n\navantage sur sauv. sagesse et charisme \n\nCommun, Goblin et 1 au choix"
	case "Yuan-Ti":
		bonus = "Cha + 2, Int + 1 \n\n30 pieds \n\nvision nocturne \n\napprend sort \"jet de poison\", peut devenir ami avec les serprents par magie,\nà partir du niveau 3 peut lancer \"suggestion\" 1 fois par jour. (utilise charisme)  \n\navantage sur jet contre effet magique \n\nimmuniser au poison \n\nCommun, Abyssal et Draconique"

	}
	return bonus
}

//randomly chooses a sous-classe in regards to the class
func sous_classe(classe string) string {
	var sous_classe = []string{}

	switch classe {
	case "Artificier":
		sous_classe = []string{"Alchimiste", "Armurier", "Artilleur", "Forgeron de bataille"}
	case "Barbare":
		sous_classe = []string{"Bête", "Berserk", "Enragé de bataille", "Guardien des ancêtres", "Guerrier du totem", "Héraut de la tempête", "Magie Sauvage", "Zélé"}
	case "Bard":
		sous_classe = []string{"Création", "Eloquence", "Esprit", "Gardien des légendes", "Glamoure", "Lame", "Murmure", "Valeur"}
	case "Cleric":
		sous_classe = []string{"Arcane", "Connaisance", "Crépuscule", "Forge", "Guerre", "Lumière", "Mort", "Nature", "Ordre", "Paix", "Tempête", "Tombe", "Tromperie", "Vie"}
	case "Combattant":
		sous_classe = []string{"Archer d'arcane", "Cavalier", "Champion", "Chevalier Echo", "Chevalier occulte", "Chevalier du dragon violet", "Chevalier des runes", "Guerrier Psy", "Maitre de bataille", "Samurai"}
	case "Démoniste":
		sous_classe = []string{"Archfée", "Céléstiale", "Génie", "Intouchable", "Immortel", "Lame maudite", "L'Ancien", "Mort-vivant", "Vilain"}
	case "Druid":
		sous_classe = []string{"Berger", "Etoiles", "Feu sauvage", "Lune", "Rêve", "Spores", "Terre"}
	case "Forestier":
		sous_classe = []string{"Chasseur", "Explorateur de l'horizon", "Guardien des drakes", "Guardien de l'essaim", "Maitre des bêtes", "Poursuivant de l'obscurité", "Sorceleur", "Voyageur du Fey"}
	case "Magicien":
		sous_classe = []string{"Abjuration", "Chronorgy", "Conjuration", "Divination", "Enchantement", "Evocation", "Guerre", "Graviturgy", "Illusion", "Nécromantie", "Scribe", "Sorcelame", "Transmutation"}
	case "Moine":
		sous_classe = []string{"Ame du soleil", "Astral", "Dragon ascendant", "Kensei", "Longue Mort", "Maitre des 4 éléments", "Main ouverte", "Ombre", "Pitié"}
	case "Paladin":
		sous_classe = []string{"Ancient", "Brise-serment", "Conquête", "Couronne", "Dévotion", "Guardien", "Gloire", "Rédemption", "Vengance"}
	case "Roublard":
		sous_classe = []string{"Assasin", "curieux", "Cerveau", "Eclaireur", "Fier-à-bras", "Lame de l'âme", "Phantome", "Trompeur arcaniste", "Voleur"}
	case "Sorcier":
		sous_classe = []string{"Ame horlogère", "Ame divine", "Draconique", "Esprit aberrant", "Ombre", "Tempête", "Sauvage"}
	}
	return choose_randomly(sous_classe)
}

//selects the bonus of the background in regards to the background
func bonuses_origine(origine string) string {
	var bonus string = ""

	switch origine {
	case "Acolyte":
		bonus = "Lire le visage, Religion \n\n2 langues au choix \n\nun symbole saint, un objet de prière, 5 bâtons d'ensens, tenu religieuse,\nvêtement commun, 15 po \n\nbienvenu dans les bâtiements de cette religion"
	case "Anthropologist":
		bonus = "Lire le visage, Religion \n\n2 langues au choix \n\nun journal, bouteille d'encre, plume pour écrire, vêtement de voyageur, 1 babiolle, 10 po \n\naprès 1 journée avec une cutlture étrangère peut communiquer avec eux (de façon rudimentaire)"
	case "Archéologiste":
		bonus = "Histoire, Survie, outils de Cartographie ou Navigation \n\n1 langue au choix \n\nune boite contenant une carte d'une ruine ou dongon, lanterne à faisceau, une pioche, vetement de voyageur, une pelle, une tente pour 2 personnes, 1 babiolle, 25 po \n\npeut déduire qui à construit une ruine et le prix d'ancienne oeuvre d'art "
	case "Artisan de guilde":
		bonus = "Lire le visage, Persuasion, 1 outil au choix \n\n1 langue au choix \n\nvos outils, lettre d'introduction de votre guilde, vetement de voyageur, 15 po \n\nbienvenu dans les bâtiments de la guilde,\nla guilde apporte son support en cas de problème légale,\ndoit payer 5 po à la guilde par mois"
	case "Athlète":
		bonus = "Acrobatics, Athlétisme, véhicule terrestre \n\n1 langue au choix \n\nun disque de bronze ou une balle de cuivre, un porte-bonheur ou trophé, vetement de voyageur, 10 po"
	case "Charlatan":
		bonus = "Mesonge, Jeux de main, kit de déguisement, kit de falsification \n\nbeau vêtement, 1 kit de déguisment, 1 outil d'arnaque (dée pipée), 15 po \n\npossède 2 identité officiel, peut falsifier des documents officiel si à déjà vue les originaux"
	case "Chasseur de prime":
		bonus = "Choisit 2 : Lire le visage, Persuasion, Discrétion, choisit 2 : kit de jeux, 1 instrument, outils de voleur \n\narmure moyenne, 20 po \n\ncontact avec vos anciens clients et les personne de même stature"
	case "Chevalier":
		bonus = "Histoire, Persuation, 1 kit de jeux \n\n1 langue au choix \n\nbeau vetement, bague avec sceau familiale, diplome d'étude, 25po \n\npossède 2 servant et 1 écuyer "
	case "Contrebandier":
		bonus = "Athlétisme, Mensonge, véhicule navale \n\nbelle veste en cuire, vetement commun, 15 po \n\npeut dormir dans des caches de contrebandier dans les grandes ville (sauf si les contrebandier locaux le refuse)"
	case "Courtier":
		bonus = "Lire le visage, Persuasion \n\n2 langues au choix \n\nbeau vêtemement, 5po \n\npeut s'imiser dans un cercle de noble comme si de rien n'était"
	case "Criminèle":
		bonus = "Mensonge, Discrétion, 1 kit de jeux, outils de voleur \n\nun pied de biche, vêtment commun sombre avec capuche, 15 po \n\npeut facilement connaitre un cercle de criminel d'une region"
	case "Divértisseur":
		bonus = "Acrobatie, Performance, kit de déguisement, 1 instrument \n\n1 instrument, le cadeau d'un fan, un costume, 15po \n\npeut se faire loger contre spectacle "
	case "Espion":
		bonus = "Mensonge, Discrétion, kit de jeux, outils de voleur \n\nun pied de biche, vêtment commun sombre avec capuche, 15 po \n\npeut facilement connaitre un cercle de criminel d'une region"
	case "Gladiateur":
		bonus = "Acrobatie, Performance, kit de déguisement, 1 instrument \n\n1 arme exotique, un cadeau de fan, costume 15po \n\npeut se faire loger en échange de spectacle "
	case "Hantée":
		bonus = "Choisit 2 : Arcane, Investigation, Religion, Survie \n\n2 langues au choix \n\n1 kit de chasseur de monstre, des vetement commun, 1 babiole, 1 pa \n\neffrai naturelement la populace"
	case "Hermit":
		bonus = "Medecine, Religion, kit d'herbaliste \n\n1 langue au choix \n\n1 porte parchemin remplie de vos recherche/notes, 1 couverture d'hiver, vetment commun, 1 kit d'herbaliste, 5po \n\nvia votre hermitage vous avez décourvert un grand secret du monde"
	case "Héritier":
		bonus = "Survie + Choisit 1 : Arcane, Histoire, Religion \\1 kit de jeux ou 1 instrument \n\n1 langue au choix \n\nton héritage, vetement de voyageur, 1 outils ou instrument, 15 po"
	case "Héro du peuple":
		bonus = "Domptage, Survie, 1 outils, véhicule terrestre \n\n1 outils, une pelle, un pot, vetement commun, 10po \n\nla populace vous logera et protegera de la loi (pas au point de ce mettre en danger)"
	case "Inspecteur":
		bonus = "Choisit 2 : Lire le visage, Investigation, Perception \\kit de déguissement, outil de voleur \n\n1 loupe, preuve d'une ancienne affaire, 1 babiole, vetement commun, 10 po \n\npeut rentrer dans les batiement officiel et est rapidement connu des autorité"
	case "Lointain voyageur":
		bonus = "Choisit 2 : Lire le visage, Perception, Athlétisme, Survie \\1 instrument ou kit de jeux \n\n1 langue au choix \n\nvetement de voyageur, 1 instrument ou kit de jeux, mauvaise carte de votre terre natale, un accesoire qui vaut 10 po, 5 po \n\npeut sans problème trouver de quoi vivre sur les chemins et à du mal à se perdre"
	case "Marchand":
		bonus = "Lire le visage, Persuasion, 1 outils ou 1 langue au choix \n\n1 langue au choix \n\n1 outil, une mule, un chariot, une lettre d'introduction de votre guilde, vetement de voyageur, 15 po \n\npeut relancer 1 jet de marchandage par jour"
	case "Marchand ratée":
		bonus = "Investigation, Persuasion, 1 outil \n\n1 langue au choix \n\n1 outils, 1 balance, beau vetement, 10 po \n\npeut demander des information à ses anciens contact"
	case "Marin":
		bonus = "Athlétisme, Perception, outils de navigateur, véhicule navale \n\n50 pieds de corde, 1 porte-bonheur, vetement commun, 10 po \n\npeut demander transport marin gratuitement"
	case "Mercenaire":
		bonus = "Choisit 2 : Mensonge, Lire le visage, Persuasion, Discrétion \\ 1 kit de jeux, véhicule terrestre \n\nuniforme de la compagnie, insigne de rang, 1 kit de jeux, 10 po \n\nconnait les compagnies mercenaire du monde et\npeut recevoir un traitement similaire à celle-ci dans une taverne ou autre"
	case "Misérable":
		bonus = "Jeux de main, Discrétion, kit de déguisement, outil de voleur \n\n1 petit couteau, carte de votre cité natale, 1 souris de compagnie,\n1 object qui vous rapelle vos parents, vetement commun, 10 po \n\npeut trouver des passages secret dans les villes"
	case "Noble":
		bonus = "Histoire, Persuasion, 1 kit de jeux \n\n1 langue au choix \n\nbeau vetement, bague avec sceau familiale, 1 diplome d'étude, 25 po \n\nrecoit un traitement de noble"
	case "Perdu du Fey":
		bonus = "Mensonge, Survie, 1 instrument \n\n1 langue au choix \n\n1 instrument, vetement de voyageur, 3 babioles, 8po \n\nles membres du fey te considère commme un des leur"
	case "Pécheur":
		bonus = "Histoire, Survie \n\n1 langue au choix \n\n1 canne à peche, 1 filet, 1 super appat, vetement de voyageur impermeable, 10po \n\navantage sur jet de peche, peut pecher pour 10 personnes par jour"
	case "Pirate":
		bonus = "Athlétisme, Perception, outils de navigateur, véhicule navale \n\n50 pieds de corde, 1 porte-bonheur, vetement commun, 10 po \n\nla populace vous craint au point de ne pas rapporter aux autorités vos petits crimes"
	case "Parieur":
		bonus = "Mensonge, Lire le visage, 1 kit de jeux \n\n1 langue au choix \n\n1 kit de jeux, 1 porte-bonheur, beau vetement, 15po \n\ninstint qui ne faille jamais (ou presque)"
	case "Sage":
		bonus = "Arcane, Histoire \n\n2 langue au choix \n\n1 bouteille d'encre noir, 1 plume, 1 petit couteau,1 lettre d'un collègue mort qui pose une question à laquelle vous n'arivez pas à répondre,vetement commun, 10 po \n\nsi ne possède pas une information, sait où la trouvée"
	case "Soldat":
		bonus = "Athlétisme, Intimidation, 1 kit de jeux, véhicule terrestre \n\ninsigne de rang, 1 trophé de guerre, 1 kit de dée fait en os + 1 kit de jeux, vetement commun, 10po \n\npeut faire appelle à ces anciens camarades "
	}
	return bonus
}
