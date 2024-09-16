use libc::c_char;
use std::ffi::CStr;

mod iqr_service;
mod types;

use iqr_service::calculate_iqr_price;
use types::InputData;

#[no_mangle]
pub extern "C" fn compute_price_statistics(json_str: *const c_char) -> f64 {
    let c_str = unsafe { CStr::from_ptr(json_str) };
    let json_str = c_str.to_str().unwrap_or("{}");
    let input_data: InputData = match serde_json::from_str(json_str) {
        Ok(data) => data,
        Err(_) => return 0.0,
    };

    let prices: Vec<f64> = input_data.prices.iter().map(|p| p.price).collect();
    let my_product_price = input_data.my_product_price;

    calculate_iqr_price(&prices, my_product_price)
}