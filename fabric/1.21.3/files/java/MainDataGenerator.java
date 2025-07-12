package %mkmod:package%;

import net.fabricmc.fabric.api.datagen.v1.DataGeneratorEntrypoint;
import net.fabricmc.fabric.api.datagen.v1.FabricDataGenerator;
:: imports ::

public class %mkmod:main%DataGenerator implements DataGeneratorEntrypoint {
    :: TOP ::
	@Override
	public void onInitializeDataGenerator(FabricDataGenerator fabricDataGenerator) {
        :: init ::
	}
	:: TAIL ::
}
