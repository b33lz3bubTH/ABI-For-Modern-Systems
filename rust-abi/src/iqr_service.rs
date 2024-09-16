use crate::types::Product;

pub fn calculate_iqr_price(prices: &[f64], _my_product_price: f64) -> f64 {
    let mut all_prices = prices.to_vec();
    // not pushing _my_product_price
    // all_prices.push(_my_product_price);
    all_prices.sort_by(|a, b| a.partial_cmp(b).unwrap_or(std::cmp::Ordering::Equal));

    let q1 = percentile(&all_prices, 25.0);
    let q3 = percentile(&all_prices, 75.0);
    let iqr = q3 - q1;

    let lower_bound = q1 - 1.5 * iqr;
    let upper_bound = q3 + 1.5 * iqr;

    let filtered_prices: Vec<f64> = all_prices.into_iter()
        .filter(|&price| price >= lower_bound && price <= upper_bound)
        .collect();

    let average = if filtered_prices.is_empty() {
        0.0
    } else {
        filtered_prices.iter().sum::<f64>() / filtered_prices.len() as f64
    };

    average
}

fn percentile(data: &[f64], percent: f64) -> f64 {
    let index = ((percent / 100.0) * (data.len() as f64)).round() as usize;
    data.get(index).copied().unwrap_or(0.0)
}