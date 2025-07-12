package %mkmod:package%;

import net.fabricmc.api.ModInitializer;
:: imports ::

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class %mkmod:main% implements ModInitializer {
	public static final String MOD_ID = "%mkmod:id%";

	// This logger is used to write text to the console and the log file.
	// It is considered best practice to use your mod id as the logger's name.
	// That way, it's clear which mod wrote info, warnings, and errors.
	public static final Logger LOGGER = LoggerFactory.getLogger(MOD_ID);
	:: TOP ::

	@Override
	public void onInitialize() {
		// This code runs as soon as Minecraft is in a mod-load-ready state.
		// However, some things (like resources) may still be uninitialized.
		// Proceed with mild caution.

		LOGGER.info("Hello Fabric world!");
		:: init ::
	}
	:: TAIL ::
}