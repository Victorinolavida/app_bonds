export function formatMonetaryValue(value: number, ){
  return new Intl.NumberFormat('en-US', {
    style: 'currency',
    maximumFractionDigits: 4,
    currency:'USD'
  }).format(value);
}

export function formatNumericValue(value: number){
  return new Intl.NumberFormat('en-US').format(value);
}