#!/bin/bash

createdb badger;

for f in migrations/*.sql;

do
  psql badger -f "$f";
done

psql badger -f seeds.sql;
