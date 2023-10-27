Project: Modern Programming Practices (CU75045V1)

Student: Rob van de Vijver 
Datum: 10 oktober 2023

Part 1 - Low(er) level programming

Reflectie:
Momenteel is het zoals beschreven staat in het artikel "Big Ball of Mud" in 1 grote functie. Ik zou de leesbaarheid en onderhoudbaarheid kunnen verbeteren door aparte functies aan te maken en deze aan te roepen. Ook kan ik gedetailleerder error handling uitleggen, zodat er precies duidelijk is wat er mis gaat. Zelf kwam ik er vaak tegen aan dat de user input een typfout bevatten, ik moet validatie gaan toepassen op data die gegeven wordt door een gebruiker om dit in de toekomst te voorkomen. Ik kan constanten gaan gebruiken zodat ik flags, namen en beschrijvingen vastleg en makkelijk vanaf 1 plek kan beheren. Ook kan ik deze makkelijk hergebruiken en voorkom "duplication code". Momenteel zit de structor van Movie in de functie en deze zou in de globale scope moeten zitten, niet in de main functie. Ik heb met de hand van AI de tip gekregen dat ik Printf format strings moet gaan gebruiken, omdat ik dan meer waardes in 1 bericht kan plaatsen. Dit zou de leesbaarheid van de code verbeteren en duplicatie verwijderen.

Na het uitvoeren van deze taken is de leesbaarheid en onderhoudbaarheid toegenomen, de tijd die nu geinvesteerd is wordt terug gewonnen op het moment van verdere ontwikkeling. De code is opnieuw gestructureerd, wat zorgt voor een verbetering in toegankelijkheid. Zelf vind ik het aanroepen van specifieke usecases wel omslachtig, een mooiere oplossing zou zijn dat de gebruiker kan interacteren met de commandline om zo tot een gesprek te komen waarin meer validatie kan plaats vinden. Ik heb ervoor gekozen dit niet toe te passen in de code, omdat dan de tests op learn zouden falen.

Part 2 - Architecture

Het is lastig te bepalen wanneer een systeem te complex wordt, daarom is ervoor gekozen om richten op de functionaliteit. Hierbij zijn aparte GO documenten aangemaakt die duidelijk omschrijven dat het gaat om routers requests. Wel is hier een discussie over te voeren over het aantal duplicatie in de code, dit zou wel beter moeten. Het is namelijk mogelijk om veel duplicatie weg te halen waardoor er minder code te lezen valt en makkelijker te begrijpen. Bepaalde componenten en relaties komen veel overeen en kunnen daarom worden samengevoegd, dit zou beheerbaarheid bevorderen. Ook het gebruik van een interface zou het gebruik van een database bevorderen. Momenteel is het zo dat ik bij iedere query de database opnieuw aanroep en open voor het uitvoeren van een query, dit is niet optimaal. Nu ik merk dat het aantal code toeneemt wordt het een must om comments te plaatsen bij functies zodat er duidelijk is wat het doel is en waarvoor het dient. Keuzes die aan het begin gemaakt zijn om deze opdracht te voltooien bleken achteraf niet handig, hierdoor is veel tijd nodig voor refactoren zodat het leesbaar wordt en minder duplicatie bevat. 


Part 3 - Concurrent programming
Momenteel is er bij de applicatie geen sprake van parallelism er gebeuren niet meerdere processen tegelijkertijd terwijl dit wel zou kunnen. Het zorgen van interfaces zou ervoor zorgen dat verschillende functionaliteiten tegelijkertijd zouden worden uitgevoerd, waardoor de doorloop tijd van functies korter wordt. Ik heb namelijk 1 loop waarin ik ieder ID opvraag en daarop een api request doe, in de toekomst zou ik het veranderen zodat er een interface is die een ID ophaald, een interface die API GET request uitvoerd, interface die het format en een status terug geeft en een interface die het JSON format in de database toevoegd. Doordat in deze architectuur niet gefocust is op beheerbaarheid kost het toevoegen van nieuwe features steeds meer tijd. Ook is het momenteel zo dat er niet duidelijk gewerkt wordt in STATUSSEN. Het is niet altijd duidelijk in welke "staat" de applicatie zich bevind omdat er 1 grote taak bezig is. Ik merk dat de infrastructuur vanaf het eerste moment goed doordacht moet zijn om op lange termijn meer tijd efficientie en begrijpbaarheid te behouden.


Part 4 - reactive programming
Bij dit hoofdstuk beoordeel ik mijn code op: complexiteit, conformiteit, veranderlijkheid en onzichtbaarheid. 
Complexiteit: De complexiteit is te overzien maar niet optimaal, er is namelijk niet gelet op code volume, het aanduiden van specifieke statussen en er is weinig controle uitgevoerd op de code. Code volume zou geoptimaliseerd kunnen worden, statussen kunnen beter gaangeduid worden door de code beter te splitsen over diverse bestanden. 

conformiteit: De code houdt zich aan conventies en standaarden volgens Svelte en Javascript, maar de code mist commentaar/documentatie om de functionaliteit aan te duiden. Ook de naamgeving van variablenen en function kunnen mogelijk duidelijker door deze te vervangen door een beschrijvende naam, waardoor de bedoelding van de code beter wordt weergeven.

veranderlijkheid: Er is gewerkt met svelte een reactive programming language, ook wordt er een API gebruikt waardoor de website dynamisch wordt bijgewerkt. Om aan te tonen dat dit zo is kun je op de revert list knop drukken, zonder dat de pagina ververst zie je dan dat de lijst is bijgewerkt.

onzichtbaarheid: voor verbetering zou ik CSS-code kunnen plaatsen op 1 centrale plek, hierdoor zou ik de onderhoudbaarheid verbeteren. Ook zou ik de cacheprestaties kunnen verbeteren door externe CSS bestanden te gebruiken (styles.css) deze worden efficiënter gecached door browsers wat de laadtijd van de pagina kan verbeteren.
