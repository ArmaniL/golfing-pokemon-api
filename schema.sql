CREATE TABLE pokemon (
    id BIGINT,

    name_english  VARCHAR,
    name_japanese VARCHAR,
    name_chinese  VARCHAR,
    name_french   VARCHAR,

    types VARCHAR[],

    base_hp          BIGINT,
    base_attack      BIGINT,
    base_defense     BIGINT,
    base_sp_attack   BIGINT,
    base_sp_defense  BIGINT,
    base_speed       BIGINT,

    species     VARCHAR,
    description VARCHAR,

    evolution_next VARCHAR[],
    evolution_prev VARCHAR[],

    height VARCHAR,
    weight VARCHAR,

    egg_groups VARCHAR[],
    abilities  VARCHAR[],
    gender     VARCHAR,

    sprite_url    VARCHAR,
    thumbnail_url VARCHAR,
    hires_url     VARCHAR
);
