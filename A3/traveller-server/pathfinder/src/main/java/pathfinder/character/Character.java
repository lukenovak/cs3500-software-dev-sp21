package pathfinder.character;

import pathfinder.town.Town;

public interface Character {

    Town getCurrentLocation();

    void setCurrentLocation(Town newLocation);

}
