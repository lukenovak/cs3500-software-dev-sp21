package pathfinder.character;

import pathfinder.town.Town;

public class CharacterImpl implements Character {

    private Town currentLocation;

    public CharacterImpl() {
        this.currentLocation= null;
    }

    @Override
    public Town getCurrentLocation() {
        return currentLocation;
    }

    @Override
    public void setCurrentLocation(Town currentLocation) {
        this.currentLocation = currentLocation;
    }
}
