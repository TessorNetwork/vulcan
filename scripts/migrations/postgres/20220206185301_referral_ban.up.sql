ALTER TABLE request
    ADD COLUMN referral_banned BOOLEAN NOT NULL DEFAULT FALSE;

UPDATE request
SET referral_banned = TRUE
WHERE address IN (
      'furya1kt4r66tgar7ed3a5fpergrna09qv3p45m6kccf',
      'furya174pu0vwklv9czuqywmvzag5ykpmcc73s2jy9de',
      'furya10g0psvuhyxzeevvayh7qjlnwfheh50k4fe6n84',
      'furya1n30nucwqukgd6y7s36nzsztpx7wp588wvz57qr',
      'furya10gpmj8n8al53p7y5cakkpddqvy4jk9htp5yu7f',
      'furya12zdrp478vm6el043sc67g53auuxgrw52vm34yp',
      'furya1zjegm8vhtd476sjwz45pma2mxug6lxaf57lzxt',
      'furya1fpaefth9uwy38sjdscev6uw7fme0padcn3qwhv',
      'furya1scpsjxn4n0vya08kd52wgqf5g5572qsfu7w7vk',
      'furya1zxrrf48xczcfcc8knwwv3kz7wdfpagfjyvxhww'
    )
