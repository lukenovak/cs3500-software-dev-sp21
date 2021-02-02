package pathfinder.traveller;

import pathfinder.character.Character;
import pathfinder.town.Town;

import java.util.List;

/**
 * Graph handler that holds the town network
 */
public interface Traveller {

    /**
     * Adds the given town to the town network map, connected to the given towns
     * @param town the town to add to the network
     * @param connectedTowns the towns that the new town is connected to
     */
    void add_to_network(Town town, List<Town> connectedTowns);

    /**
     * Places the given Character in a Town by changing a Town's occupancy state
     * @param character
     * @param town
     * @throws IllegalStateException
     */
    void place_character(Character character, Town town) throws IllegalStateException;

    /**
     * Determines whether a Character can travel to the given town without passing through
     * any other towns with Characters in them
     * @param character the Character attempting to travel
     * @param town the town that the Character is going to attempt to travel to
     * @return true if the character can travel to the town, false otherwise
     */
    Boolean can_travel_to(Character character, Town town);
}
