use pibuzzer::{Buzzer, Melody};
use rppal::gpio::Gpio;
use std::fs;

const PIN_BUZZER_NUM: u8 = 23;
const PIN_LED_NUM: u8 = 24;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let pin_led = Gpio::new()?.get(PIN_LED_NUM)?.into_output();
    let pin_buzzer = Gpio::new()?.get(PIN_BUZZER_NUM)?.into_output();
    let mut buzzer = Buzzer::new(pin_buzzer, pin_led);

    let sheet_string =
        fs::read_to_string("sheet_string").expect(&format!("Unable to read file 'sheet_string'."));

    let melody = Melody::from_str(120, &sheet_string);
    buzzer.play_melody(&melody);

    Ok(())
}
