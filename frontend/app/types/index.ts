export type User = {
  id: string;
  username: string;
};

export type Bond = {
  id: string;
  name: string;
  price: number;
  number_bonds: number;
  status: 'CREATED' | 'BONDING';
};
export type BondWithOwner = Bond & {
  owner: string;
};

export type Pagination = {
  current_page: number;
  page_size: number;
  last_page?: number;
  total_records: number;
};
