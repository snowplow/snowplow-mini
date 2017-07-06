-- Compatibility: iglu:com.example_company/example_event/jsonschema/1-0-0

CREATE TABLE atomic.com_example_company_example_event_1 (
	-- Schema of this type
	schema_vendor  varchar(128)   encode runlength not null,
	schema_name    varchar(128)   encode runlength not null,
	schema_format  varchar(128)   encode runlength not null,
	schema_version varchar(128)   encode runlength not null,
	-- Parentage of this type
	root_id        char(36)       encode raw not null,
	root_tstamp    timestamp      encode raw not null,
	ref_root       varchar(255)   encode runlength not null,
	ref_tree       varchar(1500)  encode runlength not null,
	ref_parent     varchar(255)   encode runlength not null,
	-- Properties of this type
	example_string_field       varchar(255) not null,
	example_integer_field      integer not null,
	example_numeric_field      decimal(8,2),
	example_timestamp_field    timestamp
)
DISTSTYLE KEY
-- Optimized join to atomic.events
DISTKEY (root_id)
SORTKEY (root_tstamp);