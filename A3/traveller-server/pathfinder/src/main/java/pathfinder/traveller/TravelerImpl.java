package pathfinder.traveller;

import pathfinder.character.Character;
import pathfinder.town.Town;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class TravelerImpl implements Traveller {

    private Map<Town, List<Town>> world_graph;

    public TravelerImpl() {
        this.world_graph = new HashMap<>();
    }

    @Override
    public void add_to_network(Town town, List<Town> connectedTowns) {
        this.world_graph.put(town, connectedTowns);
    }

    @Override
    public void place_character(Character character, Town town) throws IllegalStateException {
        if (town.isOccupied()) {
            throw new IllegalStateException("Town is already occupied");
        }
        character.setCurrentLocation(town);
        town.setIsOccupied(true);
    }

    @Override
    public Boolean can_travel_to(Character character, Town town) {
        Town startingTown = character.getCurrentLocation();
        return canTravelBetween(startingTown, town, new ArrayList<>());
    }

    /**
     * Depth first search to determine whether it is possible to travel between two towns
     */
    private boolean canTravelBetween(Town startingTown, Town endingTown, List<Town> alreadyVisited) {
        if (startingTown.equals(endingTown)) {
            return true;
        }
        List<Town> startingConnections = world_graph.get(startingTown);
        alreadyVisited.add(startingTown);
        for (Town town : startingConnections) {
            if (!alreadyVisited.contains(town)
                    && !town.isOccupied()
                    && canTravelBetween(town, endingTown, alreadyVisited)) {
                return true;
            }
        }
        return false;
    }
}
