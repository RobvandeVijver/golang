Project: Modern Programming Practices (CU75045V1)

Student: Rob van de Vijver 
Datum: 10 oktober 2023

Part 1 - Low(er) level programming

Reflectie:
Momenteel is het zoals beschreven staat in het artikel "Big Ball of Mud" in 1 grote functie. Ik zou de leesbaarheid en onderhoudbaarheid kunnen verbeteren door aparte functies aan te maken en deze aan te roepen. Ook kan ik gedetailleerder error handling uitleggen, zodat er precies duidelijk is wat er mis gaat. Zelf kwam ik er vaak tegen aan dat de user input een typfout bevatten, ik moet validatie gaan toepassen op data die gegeven wordt door een gebruiker om dit in de toekomst te voorkomen. Ik kan constanten gaan gebruiken zodat ik flags, namen en beschrijvingen vastleg en makkelijk vanaf 1 plek kan beheren. Ook kan ik deze makkelijk hergebruiken en voorkom "duplication code". Momenteel zit de structor van Movie in de functie en deze zou in de globale scope moeten zitten, niet in de main functie. Ik heb met de hand van AI de tip gekregen dat ik Printf format strings moet gaan gebruiken, omdat ik dan meer waardes in 1 bericht kan plaatsen. Dit zou de leesbaarheid van de code verbeteren en duplicatie verwijderen.

Na het uitvoeren van deze taken is de leesbaarheid en onderhoudbaarheid toegenomen, de tijd die nu geinvesteerd is wordt terug gewonnen op het moment van verdere ontwikkeling. De code is opnieuw gestructureerd, wat zorgt voor een verbetering in toegankelijkheid. Zelf vind ik het aanroepen van specifieke usecases wel omslachtig, een mooiere oplossing zou zijn dat de gebruiker kan interacteren met de commandline om zo tot een gesprek te komen waarin meer validatie kan plaats vinden. Ik heb ervoor gekozen dit niet toe te passen in de code, omdat dan de tests op learn zouden falen.

Part 2 - NOG NIET NAKIJKEN!