create a javascript program to generate a map of squares with the following requirements: 

I'm creating an heroic fantasy role playing game. The action takes place in a dungeon.

Dungeon: Refers to an underground or labyrinthine structure.
Room: Refers to an individual area within the dungeon.
Corridor: Refers to a passage connecting two rooms.
Door: Refers to an opening allowing passage from one room to another.
Trap: Refers to a hidden dangerous obstacle within the dungeon.
Treasure: Refers to a valuable object hidden within the dungeon.

The dungeon map is a 10 x 10 squares. Every square can be:
- a room
- a Corridor
- a Door
- a Trap

Every square can contain sometimes:
- a treasure (a number of gold pieces)
- a monster (a werewolf, an orc, a skeleton, a goblin)
- a potion for health recovery (points of health)
- a non-player character (Grignotte, Grommash)

The dungeon map has the following this json format: 

[
    {
        "squareId": "<id of the square>", 
        "squareName": "<name of the square>",
        "typeOfSquare": "<type of the sqaure>"
        "row": <row number of the square>,
        "col": <column number of the square>,
        "treasure": <number of gold pieces from 0 to 15>,
        "potion": <number of points of health from 0 to 30>,
        "monster": "<one of the possible monster>",
        "non-player character": "<one of the possible non-player character>"

    }
]

