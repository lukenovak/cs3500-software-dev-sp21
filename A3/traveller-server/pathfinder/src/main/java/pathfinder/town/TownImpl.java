package pathfinder.town;

/**
 * Basic implementation of the Town interface
 */
public class TownImpl implements Town {

    private boolean isOccupied;
    private String name;

    public TownImpl() {
        this.name = null;
    }

    @Override
    public boolean isOccupied() {
        return isOccupied;
    }

    @Override
    public void setIsOccupied(boolean isOccupied) {
        this.isOccupied = isOccupied;
    }
}
