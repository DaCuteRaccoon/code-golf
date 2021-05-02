package hole

import (
	"math/rand"
	"strings"
)

func starWarsOpeningCrawl() ([]string, string) {
	args := []string{
		"7\n37\nTurmoil has engulfed the Galactic Republic. The taxation of trade routes to outlying star systems is in dispute.\nHoping to resolve the matter with a blockade of deadly battleships, the greedy Trade Federation has stopped all shipping to the small planet of Naboo.\nWhile the Congress of the Republic endlessly debates this alarming chain of events, the Supreme Chancellor has secretly dispatched two Jedi Knights, the guardians of peace and justice in the galaxy, to settle the conflict....",
		"11\n14\nThere is unrest in the Galactic Senate. Several thousand solar systems have declared their intentions to leave the Republic.\nThis separatist movement, under the leadership of the mysterious Count Dooku, has made it difficult for the limited number of Jedi Knights to maintain peace and order in the galaxy.\nSenator Amidala, the former Queen of Naboo, is returning to the Galactic Senate to vote on the critical issue of creating an ARMY OF THE REPUBLIC to assist the overwhelmed Jedi....",
		"8\n28\nWar! The Republic is crumbling under attacks by the ruthless Sith Lord, Count Dooku. There are heroes on both sides. Evil is everywhere.\nIn a stunning move, the fiendish droid leader, General Grievous, has swept into the Republic capital and kidnapped Chancellor Palpatine, leader of the Galactic Senate.\nAs the Separatist Droid Army attempts to flee the besieged capital with their valuable hostage, two Jedi Knights lead a desperate mission to rescue the captive Chancellor....",
		"9\n23\nIt is a period of civil war. Rebel spaceships, striking from a hidden base, have won their first victory against the evil Galactic Empire.\nDuring the battle, Rebel spies managed to steal secret plans to the Empire's ultimate weapon, the DEATH STAR, an armored space station with enough power to destroy an entire planet.\nPursued by the Empire's sinister agents, Princess Leia races home aboard her starship, custodian of the stolen plans that can save her people and restore freedom to the galaxy....",
		"8\n25\nIt is a dark time for the Rebellion. Although the Death Star has been destroyed, Imperial troops have driven the Rebel forces from their hidden base and pursued them across the galaxy.\nEvading the dreaded Imperial Starfleet, a group of freedom fighters led by Luke Skywalker has established a new secret base on the remote ice world of Hoth.\nThe evil lord Darth Vader, obsessed with finding young Skywalker, has dispatched thousands of remote probes into the far reaches of space....",
		"6\n49\nLuke Skywalker has returned to his home planet of Tatooine in an attempt to rescue his friend Han Solo from the clutches of the vile gangster Jabba the Hutt.\nLittle does Luke know that the GALACTIC EMPIRE has secretly begun construction on a new armored space station even more powerful than the first dreaded Death Star.\nWhen completed, this ultimate weapon will spell certain doom for the small band of rebels struggling to restore freedom to the galaxy...",
		"7\n32\nLuke Skywalker has vanished. In his absence, the sinister FIRST ORDER has risen from the ashes of the Empire and will not rest until Skywalker, the last Jedi, has been destroyed.\nWith the support of the REPUBLIC, General Leia Organa leads a brave RESISTANCE. She is desperate to find her brother Luke and gain his help in restoring peace and justice to the galaxy.\nLeia has sent her most daring pilot on a secret mission to Jakku, where an old ally has discovered a clue to Luke's whereabouts....",
		"6\n42\nThe FIRST ORDER reigns. Having decimated the peaceful Republic, Supreme Leader Snoke now deploys his merciless legions to seize military control of the galaxy.\nOnly General Leia Organa's band of RESISTANCE fighters stand against the rising tyranny, certain that Jedi Master Luke Skywalker will return and restore a spark of hope to the fight.\nBut the Resistance has been exposed. As the First Order speeds toward the rebel base, the brave heroes mount a desperate escape....",
		"9\n17\nThe dead speak! The galaxy has heard a mysterious broadcast, a threat of REVENGE in the sinister voice of the late EMPEROR PALPATINE.\nGENERAL LEIA ORGANA dispatches secret agents to gather intelligence, while REY, the last hope of the Jedi, trains for battle against the diabolical FIRST ORDER.\nMeanwhile, Supreme Leader KYLO REN rages in search of the phantom Emperor, determined to destroy any threat to his power....",
	}

	outs := []string{
		"       Turmoil  has  engulfed  the  Galactic\n       Republic.   The   taxation  of  trade\n      routes  to  outlying star systems is in\n      dispute.\n     \n     Hoping  to  resolve  the  matter  with  a\n    blockade  of deadly battleships, the greedy\n    Trade  Federation  has stopped all shipping\n   to the small planet of Naboo.\n   \n  While  the  Congress  of the Republic endlessly\n  debates  this  alarming  chain  of  events, the\n Supreme  Chancellor  has  secretly dispatched two\n Jedi  Knights, the guardians of peace and justice\nin the galaxy, to settle the conflict....",
		"           There       is\n           unrest  in the\n          Galactic Senate.\n          Several thousand\n         solar systems have\n         declared     their\n        intentions  to leave\n        the Republic.\n       \n       This        separatist\n      movement,    under   the\n      leadership     of    the\n     mysterious   Count  Dooku,\n     has  made it difficult for\n    the  limited  number of Jedi\n    Knights  to  maintain  peace\n   and order in the galaxy.\n   \n  Senator   Amidala,   the  former\n  Queen  of Naboo, is returning to\n the Galactic Senate to vote on the\n critical issue of creating an ARMY\nOF   THE   REPUBLIC  to  assist  the\noverwhelmed Jedi....",
		"        War!    The    Republic   is\n        crumbling  under  attacks by\n       the  ruthless Sith Lord, Count\n       Dooku.  There  are  heroes  on\n      both sides. Evil is everywhere.\n      \n     In  a  stunning move, the fiendish\n     droid  leader,  General  Grievous,\n    has  swept into the Republic capital\n    and  kidnapped Chancellor Palpatine,\n   leader of the Galactic Senate.\n   \n  As the Separatist Droid Army attempts to\n  flee  the  besieged  capital  with their\n valuable  hostage, two Jedi Knights lead a\n desperate  mission  to  rescue the captive\nChancellor....",
		"         It is a period of civil\n         war.  Rebel spaceships,\n        striking  from  a  hidden\n        base,   have   won  their\n       first  victory  against the\n       evil Galactic Empire.\n      \n      During   the   battle,  Rebel\n     spies  managed  to steal secret\n     plans  to the Empire's ultimate\n    weapon,   the   DEATH   STAR,  an\n    armored space station with enough\n   power to destroy an entire planet.\n   \n  Pursued   by  the  Empire's  sinister\n  agents,   Princess  Leia  races  home\n aboard  her  starship, custodian of the\n stolen  plans  that can save her people\nand restore freedom to the galaxy....",
		"        It is a dark time for the\n        Rebellion.  Although  the\n       Death    Star    has   been\n       destroyed,  Imperial troops\n      have  driven the Rebel forces\n      from  their  hidden  base and\n     pursued them across the galaxy.\n     \n    Evading   the   dreaded  Imperial\n    Starfleet,  a  group  of  freedom\n   fighters  led by Luke Skywalker has\n   established  a  new  secret base on\n  the remote ice world of Hoth.\n  \n The  evil  lord  Darth  Vader, obsessed\n with   finding   young  Skywalker,  has\ndispatched  thousands  of  remote  probes\ninto the far reaches of space....",
		"      Luke Skywalker has returned to his home planet of\n      Tatooine  in  an attempt to rescue his friend Han\n     Solo  from  the clutches of the vile gangster Jabba\n     the Hutt.\n    \n    Little  does  Luke  know that the GALACTIC EMPIRE has\n   secretly  begun  construction  on  a  new armored space\n   station even more powerful than the first dreaded Death\n  Star.\n  \n When  completed,  this  ultimate  weapon will spell certain\n doom  for  the  small  band of rebels struggling to restore\nfreedom to the galaxy...",
		"       Luke  Skywalker has vanished. In\n       his  absence, the sinister FIRST\n      ORDER  has risen from the ashes of\n      the Empire and will not rest until\n     Skywalker,  the  last Jedi, has been\n     destroyed.\n    \n    With  the  support  of  the  REPUBLIC,\n   General   Leia   Organa  leads  a  brave\n   RESISTANCE. She is desperate to find her\n  brother   Luke   and   gain  his  help  in\n  restoring peace and justice to the galaxy.\n \n Leia  has  sent  her  most daring pilot on a\nsecret mission to Jakku, where an old ally has\ndiscovered a clue to Luke's whereabouts....",
		"      The  FIRST  ORDER reigns. Having decimated\n      the   peaceful  Republic,  Supreme  Leader\n     Snoke  now  deploys his merciless legions to\n     seize military control of the galaxy.\n    \n    Only  General Leia Organa's band of RESISTANCE\n   fighters   stand  against  the  rising  tyranny,\n   certain  that  Jedi  Master  Luke Skywalker will\n  return and restore a spark of hope to the fight.\n  \n But  the  Resistance  has been exposed. As the First\n Order speeds toward the rebel base, the brave heroes\nmount a desperate escape....",
		"         The  dead  speak!\n         The   galaxy  has\n        heard  a mysterious\n        broadcast, a threat\n       of   REVENGE  in  the\n       sinister voice of the\n      late EMPEROR PALPATINE.\n      \n     GENERAL    LEIA    ORGANA\n     dispatches  secret agents\n    to   gather   intelligence,\n    while REY, the last hope of\n   the  Jedi,  trains for battle\n   against  the diabolical FIRST\n  ORDER.\n  \n Meanwhile,  Supreme  Leader  KYLO\n REN   rages   in  search  of  the\nphantom   Emperor,   determined  to\ndestroy any threat to his power....",
	}

	rand.Shuffle(len(args), func(i, j int) {
		args[i], args[j] = args[j], args[i]
		outs[i], outs[j] = outs[j], outs[i]
	})

	return args, strings.Join(outs, "\n")
}
