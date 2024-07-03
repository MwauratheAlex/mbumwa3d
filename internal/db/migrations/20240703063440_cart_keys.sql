-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

drop table if exists cart_orders;
create table cart_orders (
  cart_id bigint not null,
  order_id bigint not null,
  primary key(cart_id, order_id),

  FOREIGN KEY (cart_id) REFERENCES carts(id) ON DELETE CASCADE,
  FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE
);
-- +goose Down
-- +goose StatementBegin
drop table cart_orders;
-- +goose StatementEnd
