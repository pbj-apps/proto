
?
!heyrenee/v1/enums/specialty.protoheyrenee.v1.enums*C
	Specialty
SPECIALTY_UNSPECIFIED 
SPECIALTY_ENDOCRINOLOGYB7Z5github.com/HeyReneeInc/proto/golang/heyrenee/v1/enumsJ?
  

  

 

 L
	
 L
:
  . Specialty specifies a healthcare speciality.



 
.
  	! The specialty is not specified.


  	

  	
.
 ! The specialty is Endocrinology.


 

 bproto3
?
#heyrenee/v1/messages/provider.protoheyrenee.v1.messages!heyrenee/v1/enums/specialty.proto"?
Provider
provider_id (	R
providerId

first_name (	R	firstName
	last_name (	RlastName#
address_lines (	RaddressLines
city (	Rcity
state (	Rstate
zip (	Rzip
phone (	Rphone:
	specialty	 (2.heyrenee.v1.enums.SpecialtyR	specialty
npi
 (	Rnpi
facility (	Rfacility
hours (	Rhours
	ribbon_id (	RribbonId-
secondary_facility (	RsecondaryFacility6
secondary_address_lines (	RsecondaryAddressLines%
secondary_city (	RsecondaryCity'
secondary_state (	RsecondaryState#
secondary_zip (	RsecondaryZip'
secondary_phone (	RsecondaryPhoneM
secondary_specialty (2.heyrenee.v1.enums.SpecialtyRsecondarySpecialty'
secondary_hours (	RsecondaryHoursB:Z8github.com/HeyReneeInc/proto/golang/heyrenee/v1/messagesJ?
  4

  

 

 O
	
 O
	
  +
<
 	 40 A Provider represents a medical care provider.



 	
&
   The ID of the Provider.


  

  	

  
4
 ' The legal first name of the Provider.


 

 	

 
3
 & The legal last name of the Provider.


 

 	

 
C
 $6 The primary address at which the Provider practices.


 


 

 

 "#
@
 3 The primary city in which the Provider practices.


 

 	

 
A
 4 The primary state in which the Provider practices.


 

 	

 
?
 2 The primary zip in which the Provider practices.


 

 	

 
D
 7 The primary phone number for contacting the Provider.


 

 	

 
?
 ,2 The primary specialty practiced by the Provider.


 

 '

 *+
T
 	G The National Provider Identifier (NPI) used to identity the Provider.


 	

 		

 	
4
 
' The primary facility of the Provider.


 


 
	

 

5
 !( The hours that the Provider practices.


 !

 !	

 !
G
 #: The ID used to identify the Provider in Ribbon's system.


 #

 #	

 #
6
 %!) The secondary facility of the Provider.


 %

 %	

 % 
E
 '/8 The secondary address at which the Provider practices.


 '


 '

 ')

 ',.
B
 )5 The secondary city in which the Provider practices.


 )

 )	

 )
C
 +6 The secondary state in which the Provider practices.


 +

 +	

 +
F
 -9 The secondary zip code in which the Provider practices.


 -

 -	

 -
F
 /9 The secondary phone number for contacting the Provider.


 /

 /	

 /
A
 174 The secondary specialty practiced by the Provider.


 1

 11

 146
?
 32 The secondary hours that the Provider practices.


 3

 3	

 3bproto3
?1
google/protobuf/timestamp.protogoogle.protobuf";
	Timestamp
seconds (Rseconds
nanos (RnanosB?
com.google.protobufBTimestampProtoPZ2google.golang.org/protobuf/types/known/timestamppb??GPB?Google.Protobuf.WellKnownTypesJ?/
 ?
?
 2? Protocol Buffers - Google's data interchange format
 Copyright 2008 Google Inc.  All rights reserved.
 https://developers.google.com/protocol-buffers/

 Redistribution and use in source and binary forms, with or without
 modification, are permitted provided that the following conditions are
 met:

     * Redistributions of source code must retain the above copyright
 notice, this list of conditions and the following disclaimer.
     * Redistributions in binary form must reproduce the above
 copyright notice, this list of conditions and the following disclaimer
 in the documentation and/or other materials provided with the
 distribution.
     * Neither the name of Google Inc. nor the names of its
 contributors may be used to endorse or promote products derived from
 this software without specific prior written permission.

 THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.


  

" ;
	
%" ;

# 
	
# 

$ I
	
$ I

% ,
	
% ,

& /
	
& /

' "
	

' "

( !
	
$( !
?
 ? ?? A Timestamp represents a point in time independent of any time zone or local
 calendar, encoded as a count of seconds and fractions of seconds at
 nanosecond resolution. The count is relative to an epoch at UTC midnight on
 January 1, 1970, in the proleptic Gregorian calendar which extends the
 Gregorian calendar backwards to year one.

 All minutes are 60 seconds long. Leap seconds are "smeared" so that no leap
 second table is needed for interpretation, using a [24-hour linear
 smear](https://developers.google.com/time/smear).

 The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By
 restricting to that range, we ensure that we can convert to and from [RFC
 3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.

 # Examples

 Example 1: Compute Timestamp from POSIX `time()`.

     Timestamp timestamp;
     timestamp.set_seconds(time(NULL));
     timestamp.set_nanos(0);

 Example 2: Compute Timestamp from POSIX `gettimeofday()`.

     struct timeval tv;
     gettimeofday(&tv, NULL);

     Timestamp timestamp;
     timestamp.set_seconds(tv.tv_sec);
     timestamp.set_nanos(tv.tv_usec * 1000);

 Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.

     FILETIME ft;
     GetSystemTimeAsFileTime(&ft);
     UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;

     // A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
     // is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
     Timestamp timestamp;
     timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
     timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));

 Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.

     long millis = System.currentTimeMillis();

     Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
         .setNanos((int) ((millis % 1000) * 1000000)).build();


 Example 5: Compute Timestamp from Java `Instant.now()`.

     Instant now = Instant.now();

     Timestamp timestamp =
         Timestamp.newBuilder().setSeconds(now.getEpochSecond())
             .setNanos(now.getNano()).build();


 Example 6: Compute Timestamp from current time in Python.

     timestamp = Timestamp()
     timestamp.GetCurrentTime()

 # JSON Mapping

 In JSON format, the Timestamp type is encoded as a string in the
 [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the
 format is "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z"
 where {year} is always expressed using four digits while {month}, {day},
 {hour}, {min}, and {sec} are zero-padded to two digits each. The fractional
 seconds, which can go up to 9 digits (i.e. up to 1 nanosecond resolution),
 are optional. The "Z" suffix indicates the timezone ("UTC"); the timezone
 is required. A proto3 JSON serializer should always use UTC (as indicated by
 "Z") when printing the Timestamp type and a proto3 JSON parser should be
 able to accept both UTC and other timezones (as indicated by an offset).

 For example, "2017-01-15T01:30:15.01Z" encodes 15.01 seconds past
 01:30 UTC on January 15, 2017.

 In JavaScript, one can convert a Date object to this format using the
 standard
 [toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
 method. In Python, a standard `datetime.datetime` object can be converted
 to this format using
 [`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with
 the time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use
 the Joda Time's [`ISODateTimeFormat.dateTime()`](
 http://www.joda.org/joda-time/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime%2D%2D
 ) to obtain a formatter capable of generating timestamps in this format.




 ?
?
  ?? Represents seconds of UTC time since Unix epoch
 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to
 9999-12-31T23:59:59Z inclusive.


  ?

  ?

  ?
?
 ?? Non-negative fractions of a second at nanosecond resolution. Negative
 second values with fractions must still have non-negative nanos values
 that count forward in time. Must be from 0 to 999,999,999
 inclusive.


 ?

 ?

 ?bproto3
?
&heyrenee/v1/messages/appointment.protoheyrenee.v1.messages#heyrenee/v1/messages/provider.protogoogle/protobuf/timestamp.proto"?
Appointment

patient_id (	R	patientId!
provider_id (	H R
providerIdK
provider_message (2.heyrenee.v1.messages.ProviderH RproviderMessage%
appointment_id (	RappointmentId.
date (2.google.protobuf.TimestampRdate#
address_lines (	RaddressLines
city (	Rcity
state (	Rstate
zip	 (	Rzip"
instructions
 (	RinstructionsB

providerB:Z8github.com/HeyReneeInc/proto/golang/heyrenee/v1/messagesJ?

  "

  

 

 O
	
 O
	
  -
	
 )
s
 
 "g An Appointment represents a Patient's scheduled event with a Provider where medical care is provided.



 

E
  8 The ID of the Patient that the Appointment belongs to.


  

  	

  
P
  B The Provider that is giving medical care during the Appointment.


  
&
  The ID of the Provider.


 


 

 
(
 " The full Provider object.


 

 

  !
)
  The ID of the Appointment.


 

 	

 
G
 %: The date and time that the Appointment is scheduled for.


 

  

 #$
A
 $4 The address where the Appointment will take place.


 


 

 

 "#
>
 1 The city where the Appointment will take place.


 

 	

 
?
 2 The state where the Appointment will take place.


 

 	

 
B
 5 The zip code where the Appointment will take place.


 

 	

 
S
 	!F Specific instructions for the Patient to follow for the Appointment.


 	!

 	!	

 	!bproto3
Ʌ
 google/protobuf/descriptor.protogoogle.protobuf"M
FileDescriptorSet8
file (2$.google.protobuf.FileDescriptorProtoRfile"?
FileDescriptorProto
name (	Rname
package (	Rpackage

dependency (	R
dependency+
public_dependency
 (RpublicDependency'
weak_dependency (RweakDependencyC
message_type (2 .google.protobuf.DescriptorProtoRmessageTypeA
	enum_type (2$.google.protobuf.EnumDescriptorProtoRenumTypeA
service (2'.google.protobuf.ServiceDescriptorProtoRserviceC
	extension (2%.google.protobuf.FieldDescriptorProtoR	extension6
options (2.google.protobuf.FileOptionsRoptionsI
source_code_info	 (2.google.protobuf.SourceCodeInfoRsourceCodeInfo
syntax (	Rsyntax"?
DescriptorProto
name (	Rname;
field (2%.google.protobuf.FieldDescriptorProtoRfieldC
	extension (2%.google.protobuf.FieldDescriptorProtoR	extensionA
nested_type (2 .google.protobuf.DescriptorProtoR
nestedTypeA
	enum_type (2$.google.protobuf.EnumDescriptorProtoRenumTypeX
extension_range (2/.google.protobuf.DescriptorProto.ExtensionRangeRextensionRangeD

oneof_decl (2%.google.protobuf.OneofDescriptorProtoR	oneofDecl9
options (2.google.protobuf.MessageOptionsRoptionsU
reserved_range	 (2..google.protobuf.DescriptorProto.ReservedRangeRreservedRange#
reserved_name
 (	RreservedNamez
ExtensionRange
start (Rstart
end (Rend@
options (2&.google.protobuf.ExtensionRangeOptionsRoptions7
ReservedRange
start (Rstart
end (Rend"|
ExtensionRangeOptionsX
uninterpreted_option? (2$.google.protobuf.UninterpretedOptionRuninterpretedOption*	?????"?
FieldDescriptorProto
name (	Rname
number (RnumberA
label (2+.google.protobuf.FieldDescriptorProto.LabelRlabel>
type (2*.google.protobuf.FieldDescriptorProto.TypeRtype
	type_name (	RtypeName
extendee (	Rextendee#
default_value (	RdefaultValue
oneof_index	 (R
oneofIndex
	json_name
 (	RjsonName7
options (2.google.protobuf.FieldOptionsRoptions'
proto3_optional (Rproto3Optional"?
Type
TYPE_DOUBLE

TYPE_FLOAT

TYPE_INT64
TYPE_UINT64

TYPE_INT32
TYPE_FIXED64
TYPE_FIXED32
	TYPE_BOOL
TYPE_STRING	

TYPE_GROUP

TYPE_MESSAGE

TYPE_BYTES
TYPE_UINT32
	TYPE_ENUM
TYPE_SFIXED32
TYPE_SFIXED64
TYPE_SINT32
TYPE_SINT64"C
Label
LABEL_OPTIONAL
LABEL_REQUIRED
LABEL_REPEATED"c
OneofDescriptorProto
name (	Rname7
options (2.google.protobuf.OneofOptionsRoptions"?
EnumDescriptorProto
name (	Rname?
value (2).google.protobuf.EnumValueDescriptorProtoRvalue6
options (2.google.protobuf.EnumOptionsRoptions]
reserved_range (26.google.protobuf.EnumDescriptorProto.EnumReservedRangeRreservedRange#
reserved_name (	RreservedName;
EnumReservedRange
start (Rstart
end (Rend"?
EnumValueDescriptorProto
name (	Rname
number (Rnumber;
options (2!.google.protobuf.EnumValueOptionsRoptions"?
ServiceDescriptorProto
name (	Rname>
method (2&.google.protobuf.MethodDescriptorProtoRmethod9
options (2.google.protobuf.ServiceOptionsRoptions"?
MethodDescriptorProto
name (	Rname

input_type (	R	inputType
output_type (	R
outputType8
options (2.google.protobuf.MethodOptionsRoptions0
client_streaming (:falseRclientStreaming0
server_streaming (:falseRserverStreaming"?	
FileOptions!
java_package (	RjavaPackage0
java_outer_classname (	RjavaOuterClassname5
java_multiple_files
 (:falseRjavaMultipleFilesD
java_generate_equals_and_hash (BRjavaGenerateEqualsAndHash:
java_string_check_utf8 (:falseRjavaStringCheckUtf8S
optimize_for	 (2).google.protobuf.FileOptions.OptimizeMode:SPEEDRoptimizeFor

go_package (	R	goPackage5
cc_generic_services (:falseRccGenericServices9
java_generic_services (:falseRjavaGenericServices5
py_generic_services (:falseRpyGenericServices7
php_generic_services* (:falseRphpGenericServices%

deprecated (:falseR
deprecated.
cc_enable_arenas (:trueRccEnableArenas*
objc_class_prefix$ (	RobjcClassPrefix)
csharp_namespace% (	RcsharpNamespace!
swift_prefix' (	RswiftPrefix(
php_class_prefix( (	RphpClassPrefix#
php_namespace) (	RphpNamespace4
php_metadata_namespace, (	RphpMetadataNamespace!
ruby_package- (	RrubyPackageX
uninterpreted_option? (2$.google.protobuf.UninterpretedOptionRuninterpretedOption":
OptimizeMode	
SPEED
	CODE_SIZE
LITE_RUNTIME*	?????J&'"?
MessageOptions<
message_set_wire_format (:falseRmessageSetWireFormatL
no_standard_descriptor_accessor (:falseRnoStandardDescriptorAccessor%

deprecated (:falseR
deprecated
	map_entry (RmapEntryX
uninterpreted_option? (2$.google.protobuf.UninterpretedOptionRuninterpretedOption*	?????JJJJ	J	
"?
FieldOptionsA
ctype (2#.google.protobuf.FieldOptions.CType:STRINGRctype
packed (RpackedG
jstype (2$.google.protobuf.FieldOptions.JSType:	JS_NORMALRjstype
lazy (:falseRlazy%

deprecated (:falseR
deprecated
weak
 (:falseRweakX
uninterpreted_option? (2$.google.protobuf.UninterpretedOptionRuninterpretedOption"/
CType

STRING 
CORD
STRING_PIECE"5
JSType
	JS_NORMAL 
	JS_STRING
	JS_NUMBER*	?????J"s
OneofOptionsX
uninterpreted_option? (2$.google.protobuf.UninterpretedOptionRuninterpretedOption*	?????"?
EnumOptions
allow_alias (R
allowAlias%

deprecated (:falseR
deprecatedX
uninterpreted_option? (2$.google.protobuf.UninterpretedOptionRuninterpretedOption*	?????J"?
EnumValueOptions%

deprecated (:falseR
deprecatedX
uninterpreted_option? (2$.google.protobuf.UninterpretedOptionRuninterpretedOption*	?????"?
ServiceOptions%

deprecated! (:falseR
deprecatedX
uninterpreted_option? (2$.google.protobuf.UninterpretedOptionRuninterpretedOption*	?????"?
MethodOptions%

deprecated! (:falseR
deprecatedq
idempotency_level" (2/.google.protobuf.MethodOptions.IdempotencyLevel:IDEMPOTENCY_UNKNOWNRidempotencyLevelX
uninterpreted_option? (2$.google.protobuf.UninterpretedOptionRuninterpretedOption"P
IdempotencyLevel
IDEMPOTENCY_UNKNOWN 
NO_SIDE_EFFECTS

IDEMPOTENT*	?????"?
UninterpretedOptionA
name (2-.google.protobuf.UninterpretedOption.NamePartRname)
identifier_value (	RidentifierValue,
positive_int_value (RpositiveIntValue,
negative_int_value (RnegativeIntValue!
double_value (RdoubleValue!
string_value (RstringValue'
aggregate_value (	RaggregateValueJ
NamePart
	name_part (	RnamePart!
is_extension (RisExtension"?
SourceCodeInfoD
location (2(.google.protobuf.SourceCodeInfo.LocationRlocation?
Location
path (BRpath
span (BRspan)
leading_comments (	RleadingComments+
trailing_comments (	RtrailingComments:
leading_detached_comments (	RleadingDetachedComments"?
GeneratedCodeInfoM

annotation (2-.google.protobuf.GeneratedCodeInfo.AnnotationR
annotationm

Annotation
path (BRpath
source_file (	R
sourceFile
begin (Rbegin
end (RendB~
com.google.protobufBDescriptorProtosHZ-google.golang.org/protobuf/types/descriptorpb??GPB?Google.Protobuf.ReflectionJ??
' ?
?
' 2? Protocol Buffers - Google's data interchange format
 Copyright 2008 Google Inc.  All rights reserved.
 https://developers.google.com/protocol-buffers/

 Redistribution and use in source and binary forms, with or without
 modification, are permitted provided that the following conditions are
 met:

     * Redistributions of source code must retain the above copyright
 notice, this list of conditions and the following disclaimer.
     * Redistributions in binary form must reproduce the above
 copyright notice, this list of conditions and the following disclaimer
 in the documentation and/or other materials provided with the
 distribution.
     * Neither the name of Google Inc. nor the names of its
 contributors may be used to endorse or promote products derived from
 this software without specific prior written permission.

 THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
2? Author: kenton@google.com (Kenton Varda)
  Based on original Protocol Buffers design by
  Sanjay Ghemawat, Jeff Dean, and others.

 The messages in this file describe the definitions found in .proto files.
 A valid .proto file can be translated directly to a FileDescriptorProto
 without any other information (e.g. without reading its imports).


) 

+ D
	
+ D

, ,
	
, ,

- 1
	
- 1

. 7
	
%. 7

/ !
	
$/ !

0 
	
0 

4 

	4 t descriptor.proto must be optimized for speed because reflection-based
 algorithms don't work during bootstrapping.

j
 8 :^ The protocol compiler can output a FileDescriptorSet containing the .proto
 files it parses.



 8

  9(

  9


  9

  9#

  9&'
/
= Z# Describes a complete .proto file.



=
9
 >", file name, relative to root of source tree


 >


 >

 >

 >
*
?" e.g. "foo", "foo.bar", etc.


?


?

?

?
4
B!' Names of files imported by this file.


B


B

B

B 
Q
D(D Indexes of the public imported files in the dependency list above.


D


D

D"

D%'
z
G&m Indexes of the weak imported files in the dependency list.
 For Google-internal migration only. Do not use.


G


G

G 

G#%
6
J,) All top-level definitions in this file.


J


J

J'

J*+

K-

K


K

K(

K+,

L.

L


L!

L")

L,-

M.

M


M

M )

M,-

	O#

	O


	O

	O

	O!"
?

U/? This field contains optional information about the original source code.
 You may safely remove this entire field without harming runtime
 functionality of the descriptors -- the information is needed only by
 development tools.



U



U


U*


U-.
]
YP The syntax of the proto file.
 The supported values are "proto2" and "proto3".


Y


Y

Y

Y
'
] } Describes a message type.



]

 ^

 ^


 ^

 ^

 ^

`*

`


`

` %

`()

a.

a


a

a )

a,-

c+

c


c

c&

c)*

d-

d


d

d(

d+,

 fk

 f


  g" Inclusive.


  g

  g

  g

  g

 h" Exclusive.


 h

 h

 h

 h

 j/

 j

 j"

 j#*

 j-.

l.

l


l

l)

l,-

n/

n


n

n *

n-.

p&

p


p

p!

p$%
?
ux? Range of reserved tag numbers. Reserved tag numbers may not be used by
 fields or extension ranges in the same message. Reserved ranges may
 not overlap.


u


 v" Inclusive.


 v

 v

 v

 v

w" Exclusive.


w

w

w

w

y,

y


y

y'

y*+
?
	|%u Reserved field names, which may not be used by fields in the same message.
 A given name may only be reserved once.


	|


	|

	|

	|"$

 ?



O
 ?:A The parser stores options it doesn't recognize here. See above.


 ?


 ?

 ?3

 ?69
Z
?M Clients can define custom options in extensions of this message. See above.


 ?

 ?

 ?
3
? ?% Describes a field within a message.


?

 ??

 ?
S
  ?C 0 is reserved for errors.
 Order is weird for historical reasons.


  ?

  ?

 ?

 ?

 ?
w
 ?g Not ZigZag encoded.  Negative numbers take 10 bytes.  Use TYPE_SINT64 if
 negative values are likely.


 ?

 ?

 ?

 ?

 ?
w
 ?g Not ZigZag encoded.  Negative numbers take 10 bytes.  Use TYPE_SINT32 if
 negative values are likely.


 ?

 ?

 ?

 ?

 ?

 ?

 ?

 ?

 ?

 ?

 ?

 ?

 ?

 ?
?
 	?? Tag-delimited aggregate.
 Group type is deprecated and not supported in proto3. However, Proto3
 implementations should still be able to parse the group wire format and
 treat group fields as unknown fields.


 	?

 	?
-
 
?" Length-delimited aggregate.


 
?

 
?
#
 ? New in version 2.


 ?

 ?

 ?

 ?

 ?

 ?

 ?

 ?

 ?

 ?

 ?

 ?

 ?

 ?
'
 ?" Uses ZigZag encoding.


 ?

 ?
'
 ?" Uses ZigZag encoding.


 ?

 ?

??

?
*
 ? 0 is reserved for errors


 ?

 ?

?

?

?

?

?

?

 ?

 ?


 ?

 ?

 ?

?

?


?

?

?

?

?


?

?

?
?
?? If type_name is set, this need not be set.  If both this and type_name
 are set, this must be one of TYPE_ENUM, TYPE_MESSAGE or TYPE_GROUP.


?


?

?

?
?
? ? For message and enum types, this is the name of the type.  If the name
 starts with a '.', it is fully-qualified.  Otherwise, C++-like scoping
 rules are used to find the type (i.e. first the nested types within this
 message are searched, then within the parent, on up to the root
 namespace).


?


?

?

?
~
?p For extensions, this is the name of the type being extended.  It is
 resolved in the same manner as type_name.


?


?

?

?
?
?$? For numeric types, contains the original text representation of the value.
 For booleans, "true" or "false".
 For strings, contains the default text contents (not escaped in any way).
 For bytes, contains the C escaped value.  All bytes >= 128 are escaped.
 TODO(kenton):  Base-64 encode?


?


?

?

?"#
?
?!v If set, gives the index of a oneof in the containing type's oneof_decl
 list.  This field is a member of that oneof.


?


?

?

? 
?
?!? JSON name of this field. The value is set by protocol compiler. If the
 user has set a "json_name" option on this field, that option's value
 will be used. Otherwise, it's deduced from the field's name by converting
 it to camelCase.


?


?

?

? 

	?$

	?


	?

	?

	?"#
?	

?%?	 If true, this is a proto3 "optional". When a proto3 field is optional, it
 tracks presence regardless of field type.

 When proto3_optional is true, this field must be belong to a oneof to
 signal to old proto3 clients that presence is tracked for this field. This
 oneof is known as a "synthetic" oneof, and this field must be its sole
 member (each proto3 optional field gets its own synthetic oneof). Synthetic
 oneofs exist in the descriptor only, and do not generate any API. Synthetic
 oneofs must be ordered after all "real" oneofs.

 For message fields, proto3_optional doesn't create any semantic change,
 since non-repeated message fields always track presence. However it still
 indicates the semantic detail of whether the user wrote "optional" or not.
 This can be useful for round-tripping the .proto file. For consistency we
 give message fields a synthetic oneof also, even though it is not required
 to track presence. This is especially important because the parser can't
 tell if a field is a message or an enum, so it must always create a
 synthetic oneof.

 Proto2 optional fields do not set this flag, because they already indicate
 optional with `LABEL_OPTIONAL`.



?



?


?


?"$
"
? ? Describes a oneof.


?

 ?

 ?


 ?

 ?

 ?

?$

?


?

?

?"#
'
? ? Describes an enum type.


?

 ?

 ?


 ?

 ?

 ?

?.

?


?#

?$)

?,-

?#

?


?

?

?!"
?
 ??? Range of reserved numeric values. Reserved values may not be used by
 entries in the same enum. Reserved ranges may not overlap.

 Note that this is distinct from DescriptorProto.ReservedRange in that it
 is inclusive such that it can appropriately represent the entire int32
 domain.


 ?


  ?" Inclusive.


  ?

  ?

  ?

  ?

 ?" Inclusive.


 ?

 ?

 ?

 ?
?
?0? Range of reserved numeric values. Reserved numeric values may not be used
 by enum values in the same enum declaration. Reserved ranges may not
 overlap.


?


?

?+

?./
l
?$^ Reserved enum value names, which may not be reused. A given name may only
 be reserved once.


?


?

?

?"#
1
? ?# Describes a value within an enum.


? 

 ?

 ?


 ?

 ?

 ?

?

?


?

?

?

?(

?


?

?#

?&'
$
? ? Describes a service.


?

 ?

 ?


 ?

 ?

 ?

?,

?


? 

?!'

?*+

?&

?


?

?!

?$%
0
	? ?" Describes a method of a service.


	?

	 ?

	 ?


	 ?

	 ?

	 ?
?
	?!? Input and output type names.  These are resolved in the same way as
 FieldDescriptorProto.type_name, but must refer to a message type.


	?


	?

	?

	? 

	?"

	?


	?

	?

	? !

	?%

	?


	?

	? 

	?#$
E
	?77 Identifies if client streams multiple client messages


	?


	?

	? 

	?#$

	?%6

	?05
E
	?77 Identifies if server streams multiple server messages


	?


	?

	? 

	?#$

	?%6

	?05
?

? ?2N ===================================================================
 Options
2? Each of the definitions above may have "options" attached.  These are
 just annotations which may cause code to be generated slightly differently
 or may contain hints for code that manipulates protocol messages.

 Clients may define custom options as extensions of the *Options messages.
 These extensions may not yet be known at parsing time, so the parser cannot
 store the values in them.  Instead it stores them in a field in the *Options
 message called uninterpreted_option. This field must have the same name
 across all *Options messages. We then use this field to populate the
 extensions when we build a descriptor, at which point all protos have been
 parsed and so all extensions are known.

 Extension numbers for custom options may be chosen as follows:
 * For options which will only be used within a single application or
   organization, or for experimental options, use field numbers 50000
   through 99999.  It is up to you to ensure that you do not use the
   same number for multiple options.
 * For options which will be published and used publicly by multiple
   independent entities, e-mail protobuf-global-extension-registry@google.com
   to reserve extension numbers. Simply provide your project name (e.g.
   Objective-C plugin) and your project website (if available) -- there's no
   need to explain how you intend to use them. Usually you only need one
   extension number. You can declare multiple options with only one extension
   number by putting them in a sub-message. See the Custom Options section of
   the docs for examples:
   https://developers.google.com/protocol-buffers/docs/proto#options
   If this turns out to be popular, a web service will be set up
   to automatically assign option numbers.



?
?

 ?#? Sets the Java package where classes generated from this .proto will be
 placed.  By default, the proto package is used, but this is often
 inappropriate because proto packages do not normally start with backwards
 domain names.



 ?



 ?


 ?


 ?!"
?

?+? Controls the name of the wrapper Java class generated for the .proto file.
 That class will always contain the .proto file's getDescriptor() method as
 well as any top-level extensions defined in the .proto file.
 If java_multiple_files is disabled, then all the other classes from the
 .proto file will be nested inside the single wrapper outer class.



?



?


?&


?)*
?

?;? If enabled, then the Java code generator will generate a separate .java
 file for each top-level message, enum, and service defined in the .proto
 file.  Thus, these types will *not* be nested inside the wrapper class
 named by java_outer_classname.  However, the wrapper class will still be
 generated to contain the file's getDescriptor() method as well as any
 top-level extensions defined in the file.



?



?


?#


?&(


?):


?49
)

?E This option does nothing.



?



?


?-


?02


?3D


?4C
?

?>? If set true, then the Java2 code generator will generate code that
 throws an exception whenever an attempt is made to assign a non-UTF-8
 byte sequence to a string field.
 Message reflection will do the same.
 However, an extension field still accepts non-UTF-8 byte sequences.
 This option has no effect on when used with the lite runtime.



?



?


?&


?)+


?,=


?7<
L

 ??< Generated classes can be optimized for speed or code size.



 ?
D

  ?"4 Generate complete code for parsing, serialization,



  ?	


  ?
G

 ? etc.
"/ Use ReflectionOps to implement these methods.



 ?


 ?
G

 ?"7 Generate code using MessageLite and the lite runtime.



 ?


 ?


?;


?



?


?$


?'(


?):


?49
?

?"? Sets the Go package where structs generated from this .proto will be
 placed. If omitted, the Go package will be derived from the following:
   - The basename of the package import path, if provided.
   - Otherwise, the package statement in the .proto file, if present.
   - Otherwise, the basename of the .proto file, without extension.



?



?


?


?!
?

?;? Should generic services be generated in each language?  "Generic" services
 are not specific to any particular RPC system.  They are generated by the
 main code generators in each language (without additional plugins).
 Generic services were the only kind of service generation supported by
 early versions of google.protobuf.

 Generic services are now considered deprecated in favor of using plugins
 that generate code specific to your particular RPC system.  Therefore,
 these default to false.  Old code which depends on generic services should
 explicitly set them to true.



?



?


?#


?&(


?):


?49


?=


?



?


?%


?(*


?+<


?6;


	?;


	?



	?


	?#


	?&(


	?):


	?49



?<



?




?



?$



?')



?*;



?5:
?

?2? Is this file deprecated?
 Depending on the target platform, this can emit Deprecated annotations
 for everything in the file, or it will be completely ignored; in the very
 least, this is a formalization for deprecating files.



?



?


?


?


? 1


?+0


?7q Enables the use of arenas for the proto messages in this file. This applies
 only to generated classes for C++.



?



?


? 


?#%


?&6


?15
?

?)? Sets the objective c class prefix which is prepended to all objective c
 generated classes from this .proto. There is no default.



?



?


?#


?&(
I

?(; Namespace for generated classes; defaults to the package.



?



?


?"


?%'
?

?$? By default Swift generators will take the proto package and CamelCase it
 replacing '.' with underscore and use that to prefix the types/symbols
 defined. When this options is provided, they will use this value instead
 to prefix the types/symbols defined.



?



?


?


?!#
~

?(p Sets the php class prefix which is prepended to all php generated classes
 from this .proto. Default is empty.



?



?


?"


?%'
?

?%? Use this option to change the namespace of php generated classes. Default
 is empty. When this option is empty, the package name will be used for
 determining the namespace.



?



?


?


?"$
?

?.? Use this option to change the namespace of php generated metadata classes.
 Default is empty. When this option is empty, the proto file name will be
 used for determining the namespace.



?



?


?(


?+-
?

?$? Use this option to change the package of ruby generated classes. Default
 is empty. When this option is not set, the package name will be used for
 determining the ruby package.



?



?


?


?!#
|

?:n The parser stores options it doesn't recognize here.
 See the documentation for the "Options" section above.



?



?


?3


?69
?

?z Clients can define custom options in extensions of this message.
 See the documentation for the "Options" section above.



 ?


 ?


 ?


	?


	 ?


	 ?


	 ?

? ?

?
?
 ?>? Set true to use the old proto1 MessageSet wire format for extensions.
 This is provided for backwards-compatibility with the MessageSet wire
 format.  You should not use this for any other reason:  It's less
 efficient, has fewer features, and is more complicated.

 The message must be defined exactly as follows:
   message Foo {
     option message_set_wire_format = true;
     extensions 4 to max;
   }
 Note that the message cannot have any defined fields; MessageSets only
 have extensions.

 All extensions of your type must be singular messages; e.g. they cannot
 be int32s, enums, or repeated messages.

 Because this is an option, the above two restrictions are not enforced by
 the protocol compiler.


 ?


 ?

 ?'

 ?*+

 ?,=

 ?7<
?
?F? Disables the generation of the standard "descriptor()" accessor, which can
 conflict with a field of the same name.  This is meant to make migration
 from proto1 easier; new code should avoid fields named "descriptor".


?


?

?/

?23

?4E

??D
?
?1? Is this message deprecated?
 Depending on the target platform, this can emit Deprecated annotations
 for the message, or it will be completely ignored; in the very least,
 this is a formalization for deprecating messages.


?


?

?

?

?0

?*/

	?

	 ?

	 ?

	 ?

	?

	?

	?

	?

	?

	?
?
?? Whether the message is an automatically generated map entry type for the
 maps field.

 For maps fields:
     map<KeyType, ValueType> map_field = 1;
 The parsed descriptor looks like:
     message MapFieldEntry {
         option map_entry = true;
         optional KeyType key = 1;
         optional ValueType value = 2;
     }
     repeated MapFieldEntry map_field = 1;

 Implementations may choose not to generate the map_entry=true message, but
 use a native map in the target language to hold the keys and values.
 The reflection APIs in such implementations still need to work as
 if the field is a repeated message field.

 NOTE: Do not set the option in .proto files. Always use the maps syntax
 instead. The option should only be implicitly set by the proto compiler
 parser.


?


?

?

?
$
	?" javalite_serializable


	?

	?

	?

	?" javanano_as_lite


	?

	?

	?
O
?:A The parser stores options it doesn't recognize here. See above.


?


?

?3

?69
Z
?M Clients can define custom options in extensions of this message. See above.


 ?

 ?

 ?

? ?

?
?
 ?.? The ctype option instructs the C++ code generator to use a different
 representation of the field than it normally would.  See the specific
 options below.  This option is not yet implemented in the open source
 release -- sorry, we'll try to include it in a future version!


 ?


 ?

 ?

 ?

 ?-

 ?&,

 ??

 ?

  ? Default mode.


  ?


  ?

 ?

 ?

 ?

 ?

 ?

 ?
?
?? The packed option can be enabled for repeated primitive fields to enable
 a more efficient representation on the wire. Rather than repeatedly
 writing the tag and type for each element, the entire array is encoded as
 a single length-delimited blob. In proto3, only explicit setting it to
 false will avoid using packed encoding.


?


?

?

?
?
?3? The jstype option determines the JavaScript type used for values of the
 field.  The option is permitted only for 64 bit integral and fixed types
 (int64, uint64, sint64, fixed64, sfixed64).  A field with jstype JS_STRING
 is represented as JavaScript string, which avoids loss of precision that
 can happen when a large value is converted to a floating point JavaScript.
 Specifying JS_NUMBER for the jstype causes the generated JavaScript code to
 use the JavaScript "number" type.  The behavior of the default option
 JS_NORMAL is implementation dependent.

 This option is an enum to permit additional types to be added, e.g.
 goog.math.Integer.


?


?

?

?

?2

?(1

??

?
'
 ? Use the default type.


 ?

 ?
)
? Use JavaScript strings.


?

?
)
? Use JavaScript numbers.


?

?
?
?+? Should this field be parsed lazily?  Lazy applies only to message-type
 fields.  It means that when the outer message is initially parsed, the
 inner message's contents will not be parsed but instead stored in encoded
 form.  The inner message will actually be parsed when it is first accessed.

 This is only a hint.  Implementations are free to choose whether to use
 eager or lazy parsing regardless of the value of this option.  However,
 setting this option true suggests that the protocol author believes that
 using lazy parsing on this field is worth the additional bookkeeping
 overhead typically needed to implement it.

 This option does not affect the public interface of any generated code;
 all method signatures remain the same.  Furthermore, thread-safety of the
 interface is not affected by this option; const methods remain safe to
 call from multiple threads concurrently, while non-const methods continue
 to require exclusive access.


 Note that implementations may choose not to check required fields within
 a lazy sub-message.  That is, calling IsInitialized() on the outer message
 may return true even if the inner message has missing required fields.
 This is necessary because otherwise the inner message would have to be
 parsed in order to perform the check, defeating the purpose of lazy
 parsing.  An implementation which chooses not to check required fields
 must be consistent about it.  That is, for any particular sub-message, the
 implementation must either *always* check its required fields, or *never*
 check its required fields, regardless of whether or not the message has
 been parsed.


?


?

?

?

?*

?$)
?
?1? Is this field deprecated?
 Depending on the target platform, this can emit Deprecated annotations
 for accessors, or it will be completely ignored; in the very least, this
 is a formalization for deprecating fields.


?


?

?

?

?0

?*/
?
?,1 For Google-internal migration only. Do not use.


?


?

?

?

?+

?%*
O
?:A The parser stores options it doesn't recognize here. See above.


?


?

?3

?69
Z
?M Clients can define custom options in extensions of this message. See above.


 ?

 ?

 ?

	?" removed jtype


	 ?

	 ?

	 ?

? ?

?
O
 ?:A The parser stores options it doesn't recognize here. See above.


 ?


 ?

 ?3

 ?69
Z
?M Clients can define custom options in extensions of this message. See above.


 ?

 ?

 ?

? ?

?
`
 ? R Set this option to true to allow mapping different tag names to the same
 value.


 ?


 ?

 ?

 ?
?
?1? Is this enum deprecated?
 Depending on the target platform, this can emit Deprecated annotations
 for the enum, or it will be completely ignored; in the very least, this
 is a formalization for deprecating enums.


?


?

?

?

?0

?*/

	?" javanano_as_lite


	 ?

	 ?

	 ?
O
?:A The parser stores options it doesn't recognize here. See above.


?


?

?3

?69
Z
?M Clients can define custom options in extensions of this message. See above.


 ?

 ?

 ?

? ?

?
?
 ?1? Is this enum value deprecated?
 Depending on the target platform, this can emit Deprecated annotations
 for the enum value, or it will be completely ignored; in the very least,
 this is a formalization for deprecating enum values.


 ?


 ?

 ?

 ?

 ?0

 ?*/
O
?:A The parser stores options it doesn't recognize here. See above.


?


?

?3

?69
Z
?M Clients can define custom options in extensions of this message. See above.


 ?

 ?

 ?

? ?

?
?
 ?2? Is this service deprecated?
 Depending on the target platform, this can emit Deprecated annotations
 for the service, or it will be completely ignored; in the very least,
 this is a formalization for deprecating services.
2? Note:  Field numbers 1 through 32 are reserved for Google's internal RPC
   framework.  We apologize for hoarding these numbers to ourselves, but
   we were already using them long before we decided to release Protocol
   Buffers.


 ?


 ?

 ?

 ?

 ? 1

 ?+0
O
?:A The parser stores options it doesn't recognize here. See above.


?


?

?3

?69
Z
?M Clients can define custom options in extensions of this message. See above.


 ?

 ?

 ?

? ?

?
?
 ?2? Is this method deprecated?
 Depending on the target platform, this can emit Deprecated annotations
 for the method, or it will be completely ignored; in the very least,
 this is a formalization for deprecating methods.
2? Note:  Field numbers 1 through 32 are reserved for Google's internal RPC
   framework.  We apologize for hoarding these numbers to ourselves, but
   we were already using them long before we decided to release Protocol
   Buffers.


 ?


 ?

 ?

 ?

 ? 1

 ?+0
?
 ??? Is this method side-effect-free (or safe in HTTP parlance), or idempotent,
 or neither? HTTP based RPC implementation may choose GET verb for safe
 methods, and PUT verb for idempotent methods instead of the default POST.


 ?

  ?

  ?

  ?
$
 ?" implies idempotent


 ?

 ?
7
 ?"' idempotent, but may have side effects


 ?

 ?

??&

?


?

?-

?02

?%

?$
O
?:A The parser stores options it doesn't recognize here. See above.


?


?

?3

?69
Z
?M Clients can define custom options in extensions of this message. See above.


 ?

 ?

 ?
?
? ?? A message representing a option the parser does not recognize. This only
 appears in options protos created by the compiler::Parser class.
 DescriptorPool resolves these when building Descriptor objects. Therefore,
 options protos in descriptor objects (e.g. returned by Descriptor::options(),
 or produced by Descriptor::CopyTo()) will never have UninterpretedOptions
 in them.


?
?
 ??? The name of the uninterpreted option.  Each string represents a segment in
 a dot-separated name.  is_extension is true iff a segment represents an
 extension (denoted with parentheses in options specs in .proto files).
 E.g.,{ ["foo", false], ["bar.baz", true], ["qux", false] } represents
 "foo.(bar.baz).qux".


 ?


  ?"

  ?

  ?

  ?

  ? !

 ?#

 ?

 ?

 ?

 ?!"

 ?

 ?


 ?

 ?

 ?
?
?'? The value of the uninterpreted option, in whatever type the tokenizer
 identified it as during parsing. Exactly one of these should be set.


?


?

?"

?%&

?)

?


?

?$

?'(

?(

?


?

?#

?&'

?#

?


?

?

?!"

?"

?


?

?

? !

?&

?


?

?!

?$%
?
? ?j Encapsulates information about the original source file from which a
 FileDescriptorProto was generated.
2` ===================================================================
 Optional source code info


?
?
 ?!? A Location identifies a piece of source code in a .proto file which
 corresponds to a particular definition.  This information is intended
 to be useful to IDEs, code indexers, documentation generators, and similar
 tools.

 For example, say we have a file like:
   message Foo {
     optional string foo = 1;
   }
 Let's look at just the field definition:
   optional string foo = 1;
   ^       ^^     ^^  ^  ^^^
   a       bc     de  f  ghi
 We have the following locations:
   span   path               represents
   [a,i)  [ 4, 0, 2, 0 ]     The whole field definition.
   [a,b)  [ 4, 0, 2, 0, 4 ]  The label (optional).
   [c,d)  [ 4, 0, 2, 0, 5 ]  The type (string).
   [e,f)  [ 4, 0, 2, 0, 1 ]  The name (foo).
   [g,h)  [ 4, 0, 2, 0, 3 ]  The number (1).

 Notes:
 - A location may refer to a repeated field itself (i.e. not to any
   particular index within it).  This is used whenever a set of elements are
   logically enclosed in a single code segment.  For example, an entire
   extend block (possibly containing multiple extension definitions) will
   have an outer location whose path refers to the "extensions" repeated
   field without an index.
 - Multiple locations may have the same path.  This happens when a single
   logical declaration is spread out across multiple places.  The most
   obvious example is the "extend" block again -- there may be multiple
   extend blocks in the same scope, each of which will have the same path.
 - A location's span is not always a subset of its parent's span.  For
   example, the "extendee" of an extension declaration appears at the
   beginning of the "extend" block and is shared by all extensions within
   the block.
 - Just because a location's span is a subset of some other location's span
   does not mean that it is a descendant.  For example, a "group" defines
   both a type and a field in a single declaration.  Thus, the locations
   corresponding to the type and field and their components will overlap.
 - Code which tries to interpret locations should probably be designed to
   ignore those that it doesn't understand, as more types of locations could
   be recorded in the future.


 ?


 ?

 ?

 ? 

 ??

 ?

?
  ?,? Identifies which part of the FileDescriptorProto was defined at this
 location.

 Each element is a field number or an index.  They form a path from
 the root FileDescriptorProto to the place where the definition.  For
 example, this path:
   [ 4, 3, 2, 7, 1 ]
 refers to:
   file.message_type(3)  // 4, 3
       .field(7)         // 2, 7
       .name()           // 1
 This is because FileDescriptorProto.message_type has field number 4:
   repeated DescriptorProto message_type = 4;
 and DescriptorProto.field has field number 2:
   repeated FieldDescriptorProto field = 2;
 and FieldDescriptorProto.name has field number 1:
   optional string name = 1;

 Thus, the above path gives the location of a field name.  If we removed
 the last element:
   [ 4, 3, 2, 7 ]
 this path refers to the whole field declaration (from the beginning
 of the label to the terminating semicolon).


  ?

  ?

  ?

  ?

  ?+

  ?*
?
 ?,? Always has exactly three or four elements: start line, start column,
 end line (optional, otherwise assumed same as start line), end column.
 These are packed into a single field for efficiency.  Note that line
 and column numbers are zero-based -- typically you will want to add
 1 to each before displaying to a user.


 ?

 ?

 ?

 ?

 ?+

 ?*
?
 ?)? If this SourceCodeInfo represents a complete declaration, these are any
 comments appearing before and after the declaration which appear to be
 attached to the declaration.

 A series of line comments appearing on consecutive lines, with no other
 tokens appearing on those lines, will be treated as a single comment.

 leading_detached_comments will keep paragraphs of comments that appear
 before (but not connected to) the current element. Each paragraph,
 separated by empty lines, will be one comment element in the repeated
 field.

 Only the comment content is provided; comment markers (e.g. //) are
 stripped out.  For block comments, leading whitespace and an asterisk
 will be stripped from the beginning of each line other than the first.
 Newlines are included in the output.

 Examples:

   optional int32 foo = 1;  // Comment attached to foo.
   // Comment attached to bar.
   optional int32 bar = 2;

   optional string baz = 3;
   // Comment attached to baz.
   // Another line attached to baz.

   // Comment attached to qux.
   //
   // Another line attached to qux.
   optional double qux = 4;

   // Detached comment for corge. This is not leading or trailing comments
   // to qux or corge because there are blank lines separating it from
   // both.

   // Detached comment for corge paragraph 2.

   optional string corge = 5;
   /* Block comment attached
    * to corge.  Leading asterisks
    * will be removed. */
   /* Block comment attached to
    * grault. */
   optional int32 grault = 6;

   // ignored detached comments.


 ?

 ?

 ?$

 ?'(

 ?*

 ?

 ?

 ?%

 ?()

 ?2

 ?

 ?

 ?-

 ?01
?
? ?? Describes the relationship between generated code and its original source
 file. A GeneratedCodeInfo message is associated with only one generated
 source file, but may contain references to different source .proto files.


?
x
 ?%j An Annotation connects some span of text in generated code to an element
 of its generating .proto file.


 ?


 ?

 ? 

 ?#$

 ??

 ?

?
  ?, Identifies the element in the original source .proto file. This field
 is formatted the same as SourceCodeInfo.Location.path.


  ?

  ?

  ?

  ?

  ?+

  ?*
O
 ?$? Identifies the filesystem path to the original source .proto.


 ?

 ?

 ?

 ?"#
w
 ?g Identifies the starting offset in bytes in the generated code
 that relates to the identified object.


 ?

 ?

 ?

 ?
?
 ?? Identifies the ending offset in bytes in the generated code that
 relates to the identified offset. The end offset should be one past
 the last relevant byte (so the length of the text = end - begin).


 ?

 ?

 ?

 ?
?
heyrenee/v1/options/auth.protoheyrenee.v1.options google/protobuf/descriptor.proto*?
FieldAuthorization
FIELD_AUTHORIZATION_NONE 
FIELD_AUTHORIZATION_PATIENT!
FIELD_AUTHORIZATION_CAREGIVER!
FIELD_AUTHORIZATION_CONCIERGE$
 FIELD_AUTHORIZATION_PRESCRIPTION$
 FIELD_AUTHORIZATION_RPM_SCHEDULE:y
field_authorization.google.protobuf.FieldOptionsІ (2'.heyrenee.v1.options.FieldAuthorizationRfieldAuthorizationB9Z7github.com/HeyReneeInc/proto/golang/heyrenee/v1/optionsJ?
  

  

 

 N
	
 N
	
  *
	
 
i
 
1^ If specified, indicates that the field requires authorization and the type of authorization.



 #


 



 
(


 
+0
=
  1 Specifies the type of authorization on a field.



 
.
  ! The field has no authorization.


  

  
Z
 "M The field is a Patient ID and access to that Patient ID must be authorized.


 

  !
^
 $Q The field is a Caregiver ID and access to that Caregiver ID must be authorized.


 

 "#
^
 $Q The field is a Concierge ID and access to that Concierge ID must be authorized.


 

 "#
d
 'W The field is a Prescription ID and access to that Prescription ID must be authorized.


 "

 %&
c
 'V The field is an RpmSchedule ID and access to that RpmSchedule ID must be authorized.


 "

 %&bproto3
?
%heyrenee/v1/appointment_service.protoheyrenee.v1&heyrenee/v1/messages/appointment.protoheyrenee/v1/options/auth.proto"?
ListAppointmentsRequest#

patient_id (	B??R	patientIdG
appointment_type (2.heyrenee.v1.AppointmentTypeRappointmentType"a
ListAppointmentsResponseE
appointments (2!.heyrenee.v1.messages.AppointmentRappointments*M
AppointmentType 
APPOINTMENT_TYPE_UNSPECIFIED 
APPOINTMENT_UPCOMING2w
AppointmentServicea
ListAppointments$.heyrenee.v1.ListAppointmentsRequest%.heyrenee.v1.ListAppointmentsResponse" B1Z/github.com/HeyReneeInc/proto/golang/heyrenee/v1J?

  )

  

 

 F
	
 F
	
  0
	
 (
?
  ? AppointmentService

 The AppointmentService provides access to a Patient's Appointments. An Appointment represents a Patient's scheduled
 event with a Provider where medical care is provided.

 URLs
  Scratch: heyrenee-v1-8jiz6fu6.uc.gateway.dev



 
Y
  UL ListAppointments returns a list of Appointments for the specified Patient.


  

  .

  9Q
]
  Q ListAppointmentsRequest is the request message for the ListAppointments method.



 
J
  b= The ID of the Patient to return Appointments for. Required.


  

  	

  

  a

  І`
@
 '3 The type of Appointments that should be returned.


 

 "

 %&
_
 !S ListAppointmentsResponse is the response message for the ListAppointments method.



 
1
  =$ The list of Appointments returned.


  


  +

  ,8

  ;<
0
 $ )$ Specifies the type of Appointment.



 $
-
  &#  No Appointment type specified.


  &

  &!"
g
 (Z The Appointment is upcoming, that is it's scheduled datetime is >= the current datetime.


 (

 (bproto3
?
google/protobuf/empty.protogoogle.protobuf"
EmptyB}
com.google.protobufB
EmptyProtoPZ.google.golang.org/protobuf/types/known/emptypb??GPB?Google.Protobuf.WellKnownTypesJ?
 3
?
 2? Protocol Buffers - Google's data interchange format
 Copyright 2008 Google Inc.  All rights reserved.
 https://developers.google.com/protocol-buffers/

 Redistribution and use in source and binary forms, with or without
 modification, are permitted provided that the following conditions are
 met:

     * Redistributions of source code must retain the above copyright
 notice, this list of conditions and the following disclaimer.
     * Redistributions in binary form must reproduce the above
 copyright notice, this list of conditions and the following disclaimer
 in the documentation and/or other materials provided with the
 distribution.
     * Neither the name of Google Inc. nor the names of its
 contributors may be used to endorse or promote products derived from
 this software without specific prior written permission.

 THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.


  

" ;
	
%" ;

# E
	
# E

$ ,
	
$ ,

% +
	
% +

& "
	

& "

' !
	
$' !

( 
	
( 
?
 3 ? A generic empty message that you can re-use to avoid defining duplicated
 empty messages in your APIs. A typical example is to use it as the request
 or the response type of an API method. For instance:

     service Foo {
       rpc Bar(google.protobuf.Empty) returns (google.protobuf.Empty);
     }

 The JSON representation for `Empty` is empty JSON object `{}`.



 3bproto3
?x
google/api/http.proto
google.api"y
Http*
rules (2.google.api.HttpRuleRrulesE
fully_decode_reserved_expansion (RfullyDecodeReservedExpansion"?
HttpRule
selector (	Rselector
get (	H Rget
put (	H Rput
post (	H Rpost
delete (	H Rdelete
patch (	H Rpatch7
custom (2.google.api.CustomHttpPatternH Rcustom
body (	Rbody#
response_body (	RresponseBodyE
additional_bindings (2.google.api.HttpRuleRadditionalBindingsB	
pattern";
CustomHttpPattern
kind (	Rkind
path (	RpathBj
com.google.apiB	HttpProtoPZAgoogle.golang.org/genproto/googleapis/api/annotations;annotations??GAPIJ?s
 ?
?
 2? Copyright 2015 Google LLC

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.


 

 
	
 

 X
	
 X

 "
	

 "

 *
	
 *

 '
	
 '

 "
	
$ "
?
  )? Defines the HTTP configuration for an API service. It contains a list of
 [HttpRule][google.api.HttpRule], each specifying the mapping of an RPC method
 to one or more HTTP REST API methods.



 
?
   ? A list of HTTP configuration rules that apply to individual API methods.

 **NOTE:** All service configuration rules follow "last one wins" order.


   


   

   

   
?
 (+? When set to true, URL path parameters will be fully URI-decoded except in
 cases of single segment matches in reserved expansion, where "%2F" will be
 left encoded.

 The default behavior is to not decode RFC 6570 reserved characters in multi
 segment matches.


 (

 (&

 ()*
?S
? ??S # gRPC Transcoding

 gRPC Transcoding is a feature for mapping between a gRPC method and one or
 more HTTP REST endpoints. It allows developers to build a single API service
 that supports both gRPC APIs and REST APIs. Many systems, including [Google
 APIs](https://github.com/googleapis/googleapis),
 [Cloud Endpoints](https://cloud.google.com/endpoints), [gRPC
 Gateway](https://github.com/grpc-ecosystem/grpc-gateway),
 and [Envoy](https://github.com/envoyproxy/envoy) proxy support this feature
 and use it for large scale production services.

 `HttpRule` defines the schema of the gRPC/REST mapping. The mapping specifies
 how different portions of the gRPC request message are mapped to the URL
 path, URL query parameters, and HTTP request body. It also controls how the
 gRPC response message is mapped to the HTTP response body. `HttpRule` is
 typically specified as an `google.api.http` annotation on the gRPC method.

 Each mapping specifies a URL path template and an HTTP method. The path
 template may refer to one or more fields in the gRPC request message, as long
 as each field is a non-repeated field with a primitive (non-message) type.
 The path template controls how fields of the request message are mapped to
 the URL path.

 Example:

     service Messaging {
       rpc GetMessage(GetMessageRequest) returns (Message) {
         option (google.api.http) = {
             get: "/v1/{name=messages/*}"
         };
       }
     }
     message GetMessageRequest {
       string name = 1; // Mapped to URL path.
     }
     message Message {
       string text = 1; // The resource content.
     }

 This enables an HTTP REST to gRPC mapping as below:

 HTTP | gRPC
 -----|-----
 `GET /v1/messages/123456`  | `GetMessage(name: "messages/123456")`

 Any fields in the request message which are not bound by the path template
 automatically become HTTP query parameters if there is no HTTP request body.
 For example:

     service Messaging {
       rpc GetMessage(GetMessageRequest) returns (Message) {
         option (google.api.http) = {
             get:"/v1/messages/{message_id}"
         };
       }
     }
     message GetMessageRequest {
       message SubMessage {
         string subfield = 1;
       }
       string message_id = 1; // Mapped to URL path.
       int64 revision = 2;    // Mapped to URL query parameter `revision`.
       SubMessage sub = 3;    // Mapped to URL query parameter `sub.subfield`.
     }

 This enables a HTTP JSON to RPC mapping as below:

 HTTP | gRPC
 -----|-----
 `GET /v1/messages/123456?revision=2&sub.subfield=foo` |
 `GetMessage(message_id: "123456" revision: 2 sub: SubMessage(subfield:
 "foo"))`

 Note that fields which are mapped to URL query parameters must have a
 primitive type or a repeated primitive type or a non-repeated message type.
 In the case of a repeated type, the parameter can be repeated in the URL
 as `...?param=A&param=B`. In the case of a message type, each field of the
 message is mapped to a separate parameter, such as
 `...?foo.a=A&foo.b=B&foo.c=C`.

 For HTTP methods that allow a request body, the `body` field
 specifies the mapping. Consider a REST update method on the
 message resource collection:

     service Messaging {
       rpc UpdateMessage(UpdateMessageRequest) returns (Message) {
         option (google.api.http) = {
           patch: "/v1/messages/{message_id}"
           body: "message"
         };
       }
     }
     message UpdateMessageRequest {
       string message_id = 1; // mapped to the URL
       Message message = 2;   // mapped to the body
     }

 The following HTTP JSON to RPC mapping is enabled, where the
 representation of the JSON in the request body is determined by
 protos JSON encoding:

 HTTP | gRPC
 -----|-----
 `PATCH /v1/messages/123456 { "text": "Hi!" }` | `UpdateMessage(message_id:
 "123456" message { text: "Hi!" })`

 The special name `*` can be used in the body mapping to define that
 every field not bound by the path template should be mapped to the
 request body.  This enables the following alternative definition of
 the update method:

     service Messaging {
       rpc UpdateMessage(Message) returns (Message) {
         option (google.api.http) = {
           patch: "/v1/messages/{message_id}"
           body: "*"
         };
       }
     }
     message Message {
       string message_id = 1;
       string text = 2;
     }


 The following HTTP JSON to RPC mapping is enabled:

 HTTP | gRPC
 -----|-----
 `PATCH /v1/messages/123456 { "text": "Hi!" }` | `UpdateMessage(message_id:
 "123456" text: "Hi!")`

 Note that when using `*` in the body mapping, it is not possible to
 have HTTP parameters, as all fields not bound by the path end in
 the body. This makes this option more rarely used in practice when
 defining REST APIs. The common usage of `*` is in custom methods
 which don't use the URL at all for transferring data.

 It is possible to define multiple HTTP methods for one RPC by using
 the `additional_bindings` option. Example:

     service Messaging {
       rpc GetMessage(GetMessageRequest) returns (Message) {
         option (google.api.http) = {
           get: "/v1/messages/{message_id}"
           additional_bindings {
             get: "/v1/users/{user_id}/messages/{message_id}"
           }
         };
       }
     }
     message GetMessageRequest {
       string message_id = 1;
       string user_id = 2;
     }

 This enables the following two alternative HTTP JSON to RPC mappings:

 HTTP | gRPC
 -----|-----
 `GET /v1/messages/123456` | `GetMessage(message_id: "123456")`
 `GET /v1/users/me/messages/123456` | `GetMessage(user_id: "me" message_id:
 "123456")`

 ## Rules for HTTP mapping

 1. Leaf request fields (recursive expansion nested messages in the request
    message) are classified into three categories:
    - Fields referred by the path template. They are passed via the URL path.
    - Fields referred by the [HttpRule.body][google.api.HttpRule.body]. They are passed via the HTTP
      request body.
    - All other fields are passed via the URL query parameters, and the
      parameter name is the field path in the request message. A repeated
      field can be represented as multiple query parameters under the same
      name.
  2. If [HttpRule.body][google.api.HttpRule.body] is "*", there is no URL query parameter, all fields
     are passed via URL path and HTTP request body.
  3. If [HttpRule.body][google.api.HttpRule.body] is omitted, there is no HTTP request body, all
     fields are passed via URL path and URL query parameters.

 ### Path template syntax

     Template = "/" Segments [ Verb ] ;
     Segments = Segment { "/" Segment } ;
     Segment  = "*" | "**" | LITERAL | Variable ;
     Variable = "{" FieldPath [ "=" Segments ] "}" ;
     FieldPath = IDENT { "." IDENT } ;
     Verb     = ":" LITERAL ;

 The syntax `*` matches a single URL path segment. The syntax `**` matches
 zero or more URL path segments, which must be the last part of the URL path
 except the `Verb`.

 The syntax `Variable` matches part of the URL path as specified by its
 template. A variable template must not contain other variables. If a variable
 matches a single path segment, its template may be omitted, e.g. `{var}`
 is equivalent to `{var=*}`.

 The syntax `LITERAL` matches literal text in the URL path. If the `LITERAL`
 contains any reserved character, such characters should be percent-encoded
 before the matching.

 If a variable contains exactly one path segment, such as `"{var}"` or
 `"{var=*}"`, when such a variable is expanded into a URL path on the client
 side, all characters except `[-_.~0-9a-zA-Z]` are percent-encoded. The
 server side does the reverse decoding. Such variables show up in the
 [Discovery
 Document](https://developers.google.com/discovery/v1/reference/apis) as
 `{var}`.

 If a variable contains multiple path segments, such as `"{var=foo/*}"`
 or `"{var=**}"`, when such a variable is expanded into a URL path on the
 client side, all characters except `[-_.~/0-9a-zA-Z]` are percent-encoded.
 The server side does the reverse decoding, except "%2F" and "%2f" are left
 unchanged. Such variables show up in the
 [Discovery
 Document](https://developers.google.com/discovery/v1/reference/apis) as
 `{+var}`.

 ## Using gRPC API Service Configuration

 gRPC API Service Configuration (service config) is a configuration language
 for configuring a gRPC service to become a user-facing product. The
 service config is simply the YAML representation of the `google.api.Service`
 proto message.

 As an alternative to annotating your proto file, you can configure gRPC
 transcoding in your service config YAML files. You do this by specifying a
 `HttpRule` that maps the gRPC method to a REST endpoint, achieving the same
 effect as the proto annotation. This can be particularly useful if you
 have a proto that is reused in multiple services. Note that any transcoding
 specified in the service config will override any matching transcoding
 configuration in the proto.

 Example:

     http:
       rules:
         # Selects a gRPC method and applies HttpRule to it.
         - selector: example.v1.Messaging.GetMessage
           get: /v1/messages/{message_id}/{sub.subfield}

 ## Special notes

 When gRPC Transcoding is used to map a gRPC to JSON REST endpoints, the
 proto to JSON conversion must follow the [proto3
 specification](https://developers.google.com/protocol-buffers/docs/proto3#json).

 While the single segment variable follows the semantics of
 [RFC 6570](https://tools.ietf.org/html/rfc6570) Section 3.2.2 Simple String
 Expansion, the multi segment variable **does not** follow RFC 6570 Section
 3.2.3 Reserved Expansion. The reason is that the Reserved Expansion
 does not expand special characters like `?` and `#`, which would lead
 to invalid URLs. As the result, gRPC Transcoding uses a custom encoding
 for multi segment variables.

 The path variables **must not** refer to any repeated or mapped field,
 because client libraries are not capable of handling such variable expansion.

 The path variables **must not** capture the leading "/" character. The reason
 is that the most common use case "{var}" does not capture the leading "/"
 character. For consistency, all path variables must share the same behavior.

 Repeated message fields must not be mapped to URL query parameters, because
 no client library can support such complicated mapping.

 If an API needs to use a JSON array for request or response body, it can map
 the request or response body to a repeated field. However, some gRPC
 Transcoding implementations may not support this feature.


?
?
 ? Selects a method to which this rule applies.

 Refer to [selector][google.api.DocumentationRule.selector] for syntax details.


 ?

 ?	

 ?
?
 ??? Determines the URL pattern is matched by this rules. This pattern can be
 used with any of the {get|put|post|delete|patch} methods. A custom method
 can be defined using the 'custom' field.


 ?
\
?N Maps to HTTP GET. Used for listing and getting information about
 resources.


?


?

?
@
?2 Maps to HTTP PUT. Used for replacing a resource.


?


?

?
X
?J Maps to HTTP POST. Used for creating a resource or performing an action.


?


?

?
B
?4 Maps to HTTP DELETE. Used for deleting a resource.


?


?

?
A
?3 Maps to HTTP PATCH. Used for updating a resource.


?


?

?
?
?!? The custom pattern is used for specifying an HTTP method that is not
 included in the `pattern` field, such as HEAD, or "*" to leave the
 HTTP method unspecified for this rule. The wild-card rule is useful
 for services that provide content to Web (HTML) clients.


?

?

? 
?
?? The name of the request field whose value is mapped to the HTTP request
 body, or `*` for mapping all request fields not captured by the path
 pattern to the HTTP body, or omitted for not having any HTTP request body.

 NOTE: the referred field must be present at the top-level of the request
 message type.


?

?	

?
?
?? Optional. The name of the response field whose value is mapped to the HTTP
 response body. When omitted, the entire response message will be used
 as the HTTP response body.

 NOTE: The referred field must be present at the top-level of the response
 message type.


?

?	

?
?
	?-? Additional HTTP bindings for the selector. Nested bindings must
 not contain an `additional_bindings` field themselves (that is,
 the nesting may only be one level deep).


	?


	?

	?'

	?*,
G
? ?9 A custom pattern is used for defining custom HTTP verb.


?
2
 ?$ The name of this custom HTTP verb.


 ?

 ?	

 ?
5
?' The path matched by this custom verb.


?

?	

?bproto3
?
google/api/annotations.proto
google.apigoogle/api/http.proto google/protobuf/descriptor.proto:K
http.google.protobuf.MethodOptions?ʼ" (2.google.api.HttpRuleRhttpBn
com.google.apiBAnnotationsProtoPZAgoogle.golang.org/genproto/googleapis/api/annotations;annotations?GAPIJ?
 
?
 2? Copyright 2015 Google LLC

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.


 
	
  
	
 *

 X
	
 X

 "
	

 "

 1
	
 1

 '
	
 '

 "
	
$ "
	
 

  See `HttpRule`.



 $


 



 


 bproto3
?
heyrenee/v1/cors.protoheyrenee.v1google/protobuf/empty.protogoogle/api/annotations.proto2\
CorsT
	Preflight.google.protobuf.Empty.google.protobuf.Empty"????B
OPTIONS/*/*B1Z/github.com/HeyReneeInc/proto/golang/heyrenee/v1J?
  

  

 

 F
	
 F
	
  %
	
 &
?
  ? CorsService

 This service should never be directly called. It exists solely because API Gateway does not support CORS headers. Any
 HTTP OPTIONS requests for CORS preflight headers are directed to this service so that headers may be set and returned.



 

  

  

  %

  0E

  

	  ?ʼ"bproto3
?
$heyrenee/v1/messages/diagnosis.protoheyrenee.v1.messages"\
	Diagnosis!
diagnosis_id (	RdiagnosisId
ICD10CM (	RICD10CM
name (	RnameB:Z8github.com/HeyReneeInc/proto/golang/heyrenee/v1/messagesJ?
  

  

 

 O
	
 O
W
  K A Diagnosis represents a diagnosable condition as specified by ICD-10-CM.



 
0
  	# The ID of the Diagnosis resource.


  	

  		

  	
?
 2 The ICD-10-CM code that specifies the Diagnosis.


 

 	

 
R
 E The human readable name of the Diagnosis as specified by ICD-10-CM.


 

 	

 bproto3
?
,heyrenee/v1/messages/patient_diagnosis.protoheyrenee.v1.messages#heyrenee/v1/messages/provider.protoheyrenee/v1/options/auth.proto$heyrenee/v1/messages/diagnosis.protogoogle/protobuf/timestamp.proto"?
PatientDiagnosis#

patient_id (	B??R	patientId#
diagnosis_id (	H RdiagnosisIdN
diagnosis_message (2.heyrenee.v1.messages.DiagnosisH RdiagnosisMessage0
patient_diagnosis_id (	RpatientDiagnosisId6
diagnosing_provider_id (	HRdiagnosingProviderId`
diagnosing_provider_message (2.heyrenee.v1.messages.ProviderHRdiagnosingProviderMessagef
patient_diagnosis_status (2,.heyrenee.v1.messages.PatientDiagnosisStatusRpatientDiagnosisStatusA
date_diagnosed (2.google.protobuf.TimestampRdateDiagnosed5
diagnosis_instructions	 (	RdiagnosisInstructions'
diagnosis_notes
 (	RdiagnosisNotesB
	diagnosisB
diagnosing_provider*?
PatientDiagnosisStatus(
$PATIENT_DIAGNOSIS_STATUS_UNSPECIFIED 
PATIENT_DIAGNOSIS_ACTIVE
PATIENT_DIAGNOSIS_INACTIVEB:Z8github.com/HeyReneeInc/proto/golang/heyrenee/v1/messagesJ?
  2

  

 

 O
	
 O
	
  -
	
 (
	
 .
	
	 )
?
  (? A PatientDiagnosis represents a relationship between a Patient and Diagnosis. The relationship indicates that
 the Patient is either currently or has in the past received the diagnosis represented by the Diagnosis resource.



 
O
  bB The ID of the Patient that has received the Diagnosis. Required.


  

  	

  

  a

  І`
=
  / The Diagnosis given to the Patient. Required.


  
F
 9 The ID of the Diagnosis given to the Patient. Required.


 


 

 
k
 $^ The Diagnosis given to the Patient. Only returned in responses, must not be set in requests.


 

 

 "#
z
 "m The ID of this PatientDiagnosis resource. Must be empty in create requests, required in all other requests.


 

 	

  !
?
 1 The Provider that made the Diagnosis. Required.


 
H
 &; The ID of the Provider that made the Diagnosis. Required.


 


 !

 $%
m
 -` The Provider that made the Diagnosis. Only returned in responses, must not be set in requests.


 

 (

 +,
<
 !6/ The status of the PatientDiagnosis. Required.


 !

 !1

 !45
E
 #/8 The date the Patient received the Diagnosis. Required.


 #

 #*

 #-.
F
 %$9 Instructions related to this Diagnosis for the Patient.


 %

 %	

 %"#
?
 	'2 Notes related to this Diagnosis for the Patient.


 	'

 	'	

 	'
3
 + 2' Specifies a PatientDiagnosis' status.



 +
;
  -+. The PatientDiagnosis' status is unspecified.


  -&

  -)*
]
 /P The Patient is actively experiencing the condition specified by the Diagnosis.


 /

 /
a
 1!T The Patient is not actively experiencing the condition specified by the Diagnosis.


 1

 1 bproto3
?!
#heyrenee/v1/diagnosis_service.protoheyrenee.v1,heyrenee/v1/messages/patient_diagnosis.protoheyrenee/v1/options/auth.proto"<
DiagnosisSuggestRequest!
partial_text (	RpartialText"q
DiagnosisSuggestResponseU
diagnosis_suggestions (2 .heyrenee.v1.DiagnosisSuggestionRdiagnosisSuggestions"[
DiagnosisSuggestion!
diagnosis_id (	RdiagnosisId!
display_name (	RdisplayName"t
CreatePatientDiagnosisRequestS
patient_diagnosis (2&.heyrenee.v1.messages.PatientDiagnosisRpatientDiagnosis"t
UpdatePatientDiagnosisRequestS
patient_diagnosis (2&.heyrenee.v1.messages.PatientDiagnosisRpatientDiagnosis"?
ListPatientDiagnosesRequest#

patient_id (	B??R	patientIdf
patient_diagnosis_status (2,.heyrenee.v1.messages.PatientDiagnosisStatusRpatientDiagnosisStatus"s
ListPatientDiagnosesResponseS
patient_diagnoses (2&.heyrenee.v1.messages.PatientDiagnosisRpatientDiagnoses2?
DiagnosisServicea
DiagnosisSuggest$.heyrenee.v1.DiagnosisSuggestRequest%.heyrenee.v1.DiagnosisSuggestResponse" n
CreatePatientDiagnosis*.heyrenee.v1.CreatePatientDiagnosisRequest&.heyrenee.v1.messages.PatientDiagnosis" n
UpdatePatientDiagnosis*.heyrenee.v1.UpdatePatientDiagnosisRequest&.heyrenee.v1.messages.PatientDiagnosis" m
ListPatientDiagnoses(.heyrenee.v1.ListPatientDiagnosesRequest).heyrenee.v1.ListPatientDiagnosesResponse" B1Z/github.com/HeyReneeInc/proto/golang/heyrenee/v1J?
  M

  

 

 F
	
 F
	
  6
	
 (
?
  ? DiagnosisService

 The DiagnosisService provides operations on Diagnosis resources along with managing PatientDiagnoses which represent
 relationships between Patient users and Diagnoses.

 URLs
  Scratch: heyrenee-v1-8jiz6fu6.uc.gateway.dev



 
?
  U? DiagnosisSuggest returns a list of DiagnosisSuggestions based on a partial search string. This should only be
 used for autocomplete like features and not as a full fledged Diagnosis search method.

 TODO(mitch): Use bi-directional streaming once client streaming supported by grpc-web
 https://github.com/grpc/grpc-web/issues/24


  

  .

  9Q
K
 n> CreatePatientDiagnosis creates a specified PatientDiagnosis.


 

 :

 Ej
K
 n> UpdatePatientDiagnosis creates a specified PatientDiagnosis.


 

 :

 Ej
W
 aJ ListPatientDiagnoses lists the PatientDiagnoses for a specified Patient.


 

 6

 A]
3
 ! $' Request message for DiagnosisSuggest.



 !
S
  #F The partial text be used to generate DiagnosisSuggestions. Required.


  #

  #	

  #
4
' *( Response message for DiagnosisSuggest.



' 
0
 )9# The list of DiagnosisSuggestions.


 )


 )

 )4

 )78
?
. 3? A DiagnosisSuggestion represents a single Diagnosis suggested as an autocompletion result for the partial text in a
 DiagnosisSuggestRequest.



.
1
 0$ The ID of the suggested Diagnosis.


 0

 0	

 0
i
2\ The name of the Diagnosis that should be displayed to the user in the list of suggestions.


2

2	

2
9
6 9- Request message for CreatePatientDiagnosis.



6%
8
 8>+ The PatientDiagnosis to create. Required.


 8'

 8(9

 8<=
9
< ?- Request message for UpdatePatientDiagnosis.



<%
8
 >>+ The PatientDiagnosis to update. Required.


 >'

 >(9

 ><=
7
B G+ Request message for ListPatientDiagnoses.



B#
J
 Db= The ID of the Patient to return Appointments for. Required.


 D

 D	

 D

 Da

 ІD`
?
FKu The status of the PatientDiagnoses. If specified, only PatientDiagnoses with the specified status will be returned.


F-

F.F

FIJ
8
J M, Response message for ListPatientDiagnoses.



J$
,
 LG The list of PatientDiagnoses.


 L


 L0

 L1B

 LEFbproto3
?-
$heyrenee/v1/messages/insurance.protoheyrenee.v1.messagesheyrenee/v1/options/auth.proto"?
	Insurance#

patient_id (	B??R	patientId!
insurance_id (	RinsuranceId
insurer (	RinsurerJ
insurance_type (2#.heyrenee.v1.messages.InsuranceTypeRinsuranceTypeP
insurance_status (2%.heyrenee.v1.messages.InsuranceStatusRinsuranceStatus*
policy_owner_name (	RpolicyOwnerName;
policy_owner_address_lines (	RpolicyOwnerAddressLines*
policy_owner_city (	RpolicyOwnerCity,
policy_owner_state	 (	RpolicyOwnerState(
policy_owner_zip
 (	RpolicyOwnerZip,
policy_owner_phone (	RpolicyOwnerPhone!
group_number (	RgroupNumber#
policy_number (	RpolicyNumber(
rx_policy_number (	RrxPolicyNumber0
claims_address_lines (	RclaimsAddressLines
claims_city (	R
claimsCity!
claims_state (	RclaimsState

claims_zip (	R	claimsZip!
claims_phone (	RclaimsPhone5
rx_claims_address_lines (	RrxClaimsAddressLines$
rx_claims_city (	RrxClaimsCity&
rx_claims_state (	RrxClaimsState"
rx_claims_zip (	RrxClaimsZip&
rx_claims_phone (	RrxClaimsPhone*?
InsuranceType
INSURANCE_TYPE_UNSPECIFIED 
INSURANCE_TYPE_MEDICAID
INSURANCE_TYPE_MEDICARE%
!INSURANCE_TYPE_MEDICARE_ADVANTAGE#
INSURANCE_TYPE_VETERANS_AFFAIRS!
INSURANCE_TYPE_EMPLOYER_BASED
INSURANCE_TYPE_PRIVATE*o
InsuranceStatus 
INSURANCE_STATUS_UNSPECIFIED 
INSURANCE_STATUS_ACTIVE
INSURANCE_STATUS_INACTIVEB:Z8github.com/HeyReneeInc/proto/golang/heyrenee/v1/messagesJ? 
  V

  

 

 O
	
 O
	
  (
n
 	 :b Insurance represents an insurance policy that is currently or has in the past covered a Patient.



 	
I
  b< The ID of the Patient that the Insurance covers. Required.


  

  	

  

  a

  І`
r
 e The ID of the Insurance. Must not be provided for Create requests, required for all other requests.


 

 	

 
:
 - The name of the insuring company. Required.


 

 	

 
H
 #; The type of Insurance. Required, must not be unspecified.


 

 

 !"
N
 'A The status of the Insurance. Required, must not be unspecified.


 

 "

 %&
I
 < The name of the policy owner or primary insured. Required.


 

 	

 
v
 1i The address lines on file with the insurance company for the policy owner or primary insured. Required.


 


 

 ,

 /0
m
 ` The city on file with the insurance company for the policy owner or primary insured. Required.


 

 	

 
n
  a The state on file with the insurance company for the policy owner or primary insured. Required.


 

 	

 
q
 	d The zip code on file with the insurance company for the policy owner or primary insured. Required.


 	

 		

 	
u
 
!h The phone number on file with the insurance company for the policy owner or primary insured. Required.


 


 
	

 
 
B
 !5 The group number of the insurance policy. Required.


 !

 !	

 !
C
 #6 The policy number of the insurance policy. Required.


 #

 #	

 #
<
 %/ The Rx policy number of the insurance policy.


 %

 %	

 %
=
 ',0 The address lines for filing claims. Required.


 '


 '

 '&

 ')+
4
 )' The city for filing claims. Required.


 )

 )	

 )
5
 +( The state for filing claims. Required.


 +

 +	

 +
8
 -+ The zip code for filing claims. Required.


 -

 -	

 -
<
 // The phone number for filing claims. Required.


 /

 /	

 /
6
 1/) The address lines for filing Rx claims.


 1


 1

 1)

 1,.
-
 3  The city for filing Rx claims.


 3

 3	

 3
.
 5! The state for filing Rx claims.


 5

 5	

 5
1
 7$ The zip code for filing Rx claims.


 7

 7	

 7
5
 9( The phone number for filing Rx claims.


 9

 9	

 9
h
 = L\ InsuranceType specifies the type of insurance policy represented by an Insurance resource.



 =
3
  ?!& The Insurance type is not specified.


  ?

  ? 
2
 A% The Insurance is a medicaid policy.


 A

 A
2
 C% The Insurance is a medicare policy.


 C

 C
<
 E(/ The Insurance is a medicare advantage policy.


 E#

 E&'
:
 G%- The Insurance is a Veterans Affairs policy.


 G!

 G#$
9
 I$, The Insurance is an employer based policy.


 I

 I"#
@
 K3 The Insurance is a supplied by private insurance.


 K

 K
q
O Ve InsuranceStatus represents the status of the insurance policy represented by an Insurance resource.



O
5
 Q#( The Insurance status is not specified.


 Q

 Q!"
C
S6 The Insurance is currently and covering the Patient.


S

S
L
U ? The Insurance is inactive and no longer covering the Patient.


U

Ubproto3
?
#heyrenee/v1/insurance_service.protoheyrenee.v1$heyrenee/v1/messages/insurance.protoheyrenee/v1/options/auth.proto"W
CreateInsuranceRequest=
	insurance (2.heyrenee.v1.messages.InsuranceR	insurance"W
UpdateInsuranceRequest=
	insurance (2.heyrenee.v1.messages.InsuranceR	insurance"?
ListInsuranceRequest#

patient_id (	B??R	patientIdP
insurance_status (2%.heyrenee.v1.messages.InsuranceStatusRinsuranceStatus"V
ListInsuranceResponse=
	insurance (2.heyrenee.v1.messages.InsuranceR	insurance2?
InsuranceServiceY
CreateInsurance#.heyrenee.v1.CreateInsuranceRequest.heyrenee.v1.messages.Insurance" Y
UpdateInsurance#.heyrenee.v1.UpdateInsuranceRequest.heyrenee.v1.messages.Insurance" X
ListInsurance!.heyrenee.v1.ListInsuranceRequest".heyrenee.v1.ListInsuranceResponse" B1Z/github.com/HeyReneeInc/proto/golang/heyrenee/v1J?
  0

  

 

 F
	
 F
	
  .
	
 (
?
  ? InsuranceService

 The InsuranceService provides operations on insurance related resources that belong to specific Patients.

 URLs
  Scratch: heyrenee-v1-8jiz6fu6.uc.gateway.dev



 
E
  Y8 CreateInsurance creates a provided Insurance resource.


  

  ,

  7U
E
 Y8 UpdateInsurance updates a provided Insurance resource.


 

 ,

 7U
O
 LB ListInsurance lists Insurance resources for a specified Patient.


 

 (

 3H
2
  & Request message for CreateInsurance.



 
:
  /- The Insurance resource to create. Required.


   

  !*

  -.
2
 "& Request message for UpdateInsurance.




:
 !/- The Insurance resource to update. Required.


 ! 

 !!*

 !-.
0
% *$ Request message for ListInsurance.



%
G
 'b: The ID of the Patient to return Insurance for. Required.


 '

 '	

 '

 'a

 І'`
?
)<2 The status of Insurance that should be returned.


)&

)'7

):;
1
- 0% Response message for ListInsurance.



-
%
 /8 The list of Insurance.


 /


 /)

 /*3

 /67bproto3
?
*heyrenee/v1/messages/medication_dose.protoheyrenee.v1.messagesgoogle/protobuf/timestamp.proto"?
MedicationDose

patient_id (	R	patientId#
medication_id (	RmedicationId'
prescription_id (	RprescriptionId,
medication_dose_id (	RmedicationDoseId9

time_taken (2.google.protobuf.TimestampR	timeTakenB:Z8github.com/HeyReneeInc/proto/golang/heyrenee/v1/messagesJ?
  

  

 

 O
	
 O
	
  )
x
 	 l A MedicationDose represents a single dose of Medication taken by a Patient as specified in a Prescription.



 	
A
  4 The ID of the Patient who took the MedicationDose.


  

  	

  
7
 * The ID of the Medication that was taken.


 

 	

 
X
 K The ID of the Prescription that prescribed the Medication that was taken.


 

 	

 
,
   The ID of the MedicationDose.


 

 	

 
>
 +1 The time at which the MedicationDose was taken.


 

 &

 )*bproto3
?%
google/protobuf/duration.protogoogle.protobuf":
Duration
seconds (Rseconds
nanos (RnanosB?
com.google.protobufBDurationProtoPZ1google.golang.org/protobuf/types/known/durationpb??GPB?Google.Protobuf.WellKnownTypesJ?#
 s
?
 2? Protocol Buffers - Google's data interchange format
 Copyright 2008 Google Inc.  All rights reserved.
 https://developers.google.com/protocol-buffers/

 Redistribution and use in source and binary forms, with or without
 modification, are permitted provided that the following conditions are
 met:

     * Redistributions of source code must retain the above copyright
 notice, this list of conditions and the following disclaimer.
     * Redistributions in binary form must reproduce the above
 copyright notice, this list of conditions and the following disclaimer
 in the documentation and/or other materials provided with the
 distribution.
     * Neither the name of Google Inc. nor the names of its
 contributors may be used to endorse or promote products derived from
 this software without specific prior written permission.

 THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.


  

" ;
	
%" ;

# 
	
# 

$ H
	
$ H

% ,
	
% ,

& .
	
& .

' "
	

' "

( !
	
$( !
?
 f s? A Duration represents a signed, fixed-length span of time represented
 as a count of seconds and fractions of seconds at nanosecond
 resolution. It is independent of any calendar and concepts like "day"
 or "month". It is related to Timestamp in that the difference between
 two Timestamp values is a Duration and it can be added or subtracted
 from a Timestamp. Range is approximately +-10,000 years.

 # Examples

 Example 1: Compute Duration from two Timestamps in pseudo code.

     Timestamp start = ...;
     Timestamp end = ...;
     Duration duration = ...;

     duration.seconds = end.seconds - start.seconds;
     duration.nanos = end.nanos - start.nanos;

     if (duration.seconds < 0 && duration.nanos > 0) {
       duration.seconds += 1;
       duration.nanos -= 1000000000;
     } else if (duration.seconds > 0 && duration.nanos < 0) {
       duration.seconds -= 1;
       duration.nanos += 1000000000;
     }

 Example 2: Compute Timestamp from Timestamp + Duration in pseudo code.

     Timestamp start = ...;
     Duration duration = ...;
     Timestamp end = ...;

     end.seconds = start.seconds + duration.seconds;
     end.nanos = start.nanos + duration.nanos;

     if (end.nanos < 0) {
       end.seconds -= 1;
       end.nanos += 1000000000;
     } else if (end.nanos >= 1000000000) {
       end.seconds += 1;
       end.nanos -= 1000000000;
     }

 Example 3: Compute Duration from datetime.timedelta in Python.

     td = datetime.timedelta(days=3, minutes=10)
     duration = Duration()
     duration.FromTimedelta(td)

 # JSON Mapping

 In JSON format, the Duration type is encoded as a string rather than an
 object, where the string ends in the suffix "s" (indicating seconds) and
 is preceded by the number of seconds, with nanoseconds expressed as
 fractional seconds. For example, 3 seconds with 0 nanoseconds should be
 encoded in JSON format as "3s", while 3 seconds and 1 nanosecond should
 be expressed in JSON format as "3.000000001s", and 3 seconds and 1
 microsecond should be expressed in JSON format as "3.000001s".





 f
?
  j? Signed seconds of the span of time. Must be from -315,576,000,000
 to +315,576,000,000 inclusive. Note: these bounds are computed from:
 60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years


  j

  j

  j
?
 r? Signed fractions of a second at nanosecond resolution of the span
 of time. Durations less than one second are represented with a 0
 `seconds` field and a positive or negative `nanos` field. For durations
 of one second or more, a non-zero value for the `nanos` field must be
 of the same sign as the `seconds` field. Must be from -999,999,999
 to +999,999,999 inclusive.


 r

 r

 rbproto3
?
%heyrenee/v1/messages/medication.protoheyrenee.v1.messages"?

Medication#
medication_id (	RmedicationId
rxcui (	Rrxcui#
generic_rxcui (	RgenericRxcui
	term_type (	RtermType
	full_name (	RfullName)
rx_norm_dose_form (	RrxNormDoseForm*
full_generic_name (	RfullGenericName

brand_name (	R	brandName!
display_name	 (	RdisplayName
route
 (	Rroute"
new_dose_form (	RnewDoseForm
strength (	Rstrength!
suppress_for (	RsuppressFor0
display_name_synonym (	RdisplayNameSynonym

sxdg_rxcui (	R	sxdgRxcui$
sxdg_term_type (	RsxdgTermType
	sxdg_name (	RsxdgName+
prescribable_name (	RprescribableNameB:Z8github.com/HeyReneeInc/proto/golang/heyrenee/v1/messagesJ?
  4

  

 

 O
	
 O
F
  4: A Medication represents a distinct and unique medication



 
)
  	 The ID of the Medication..


  	

  		

  	
?
 ? The RxNorm unique identifier for the clinical drug, which can be one of the following term types:
    - Semantic Clinical Drug (SCD) e.g. azithromycin 250 MG Oral Capsule.
    - Semantic Branded Drug (SBD) e.g. azithromycin 250 MG Oral Capsule [Zithromax].
    - Generic Pack (GPCK) e.g. {6 (azithromycin 250 MG Oral Tablet) } Pack
    - Branded Pack (BPCK) e.g. {6 (azithromycin 250 MG Oral Tablet [Zithromax]) } Pack [Z-PAKS].


 

 	

 
7
 * The corresponding generic clinical drug.


 

 	

 
'
  The term type in RxNorm.


 

 	

 
9
 , The full RxNorm name of the clinical drug.


 

 	

 
D
 7 Dose form and intended route information from RxNorm.


 

 	

 
8
 + The generic part of the full RxNorm name.


 

 	

 
;
 . The brand name part of the full RxNorm name.


 

 	

 
d
 W Drug name (either generic or brand name) and intended route e.g. INDERAL (Oral-pill).


 

 	

 
9
 	, Intended route derived from RXN_DOSE_FORM.


 	

 		

 	
4
 
!' Dose form derived from RXN_DOSE_FORM.


 
!

 
!	

 
!
E
 #8 Strength information parsed from the RxNorm full name.


 #

 #	

 #
?
 '? To flag drug names deemed not likely to be useful for data entry. For example, long generic drug names with multiple
 ingredients (e.g. Bacitracin/Hydrocortisone/Neomycin/Polymyxin B) are suppressed because they are almost always
 prescribed by their brand names (e.g. CORTISPORIN OINTMENT). Any non-empty value means that a row should be suppressed.


 '

 '	

 '
b
 )#U Commonly used synonyms or abbreviations for the drug e.g. MOM for Milk of Magnesia.


 )

 )	

 ) "
?
 -? The RxNorm unique identifier for the entity represented by the DISPLAY_NAME (drug + intended route). The
 corresponding TTY (term type) in RxNorm is either SCDG (Semantic Clinical Doseform Group) or SBDG (Semantic Branded
 Doseform Group), as represented in the field SXDG_TTY.


 -

 -	

 -
=
 /0 The RxNorm TTY of the semantic doseform group.


 /

 /	

 /
>
 11 The RxNorm name of the semantic doseform group.


 1

 1	

 1
,
 3  The RxNorm prescribable name.


 3

 3	

 3bproto3
?&
'heyrenee/v1/messages/prescription.protoheyrenee.v1.messagesgoogle/protobuf/duration.protogoogle/protobuf/timestamp.proto%heyrenee/v1/messages/medication.proto#heyrenee/v1/messages/provider.protoheyrenee/v1/options/auth.proto"?
Prescription#

patient_id (	B??R	patientId%
medication_id (	H RmedicationIdQ
medication_message (2 .heyrenee.v1.messages.MedicationH RmedicationMessage'
prescription_id (	RprescriptionId!
provider_id (	HR
providerIdK
provider_message (2.heyrenee.v1.messages.ProviderHRproviderMessage3
provider_instructions (	RproviderInstructions1
patient_instructions (	RpatientInstructions!
refill_count	 (RrefillCountD
refill_frequency
 (2.google.protobuf.DurationRrefillFrequencyY
prescription_status (2(.heyrenee.v1.messages.PrescriptionStatusRprescriptionStatus_
first_medication_regimen_start (2.google.protobuf.TimestampRfirstMedicationRegimenStartY
medication_regimen_duration (2.google.protobuf.DurationRmedicationRegimenDuration?
medication_doses_per_regimen (RmedicationDosesPerRegimenx
,medication_dose_durations_from_regimen_start (2.google.protobuf.DurationR'medicationDoseDurationsFromRegimenStart=
date_written (2.google.protobuf.TimestampRdateWrittenB

medicationB

provider*m
PrescriptionStatus#
PRESCRIPTION_STATUS_UNSPECIFIED 
PRESCRIPTION_ACTIVE
PRESCRIPTION_INACTIVEB:Z8github.com/HeyReneeInc/proto/golang/heyrenee/v1/messagesJ?
  B

  

 

 O
	
 O
	
  (
	
 )
	
 /
	
	 -
	

 (
?
  8? A Prescription represents a Patient's Prescription for Medication written by a Provider. Multiple written prescriptions
 for the same Medication will have separate Prescription resources.



 
T
  bG The ID of the Patient that the Prescription is written for. Required.


  

  	

  

  a

  І`
N
  @ The Medication that the Prescription is written for. Required.


  
W
 J The ID of the Medication that the Prescription is written for. Required.


 


 

 
|
 &o The Medication that the Prescription is written for. Only returned in responses, must not be set in requests.


 

 !

 $%
i
 \ The ID of the Prescription. Must be empty in create requests, required in update requests.


 

 	

 
C
  5 The Provider that wrote the Prescription. Required.


 
L
 ? The ID of the Provider that wrote the Prescription. Required.


 


 

 
q
 "d The Provider that wrote the Prescription. Only returned in responses, must not be set in requests.


 

 

  !
f
 "#Y Instructions supplied by the Provider on how the prescribed Medication should be taken.


 "

 "	

 "!"
e
 $"X Instructions supplied by the Patient on how the prescribed Medication should be taken.


 $

 $	

 $ !
_
 &R The total number of refills included in the Prescription. Required, must be > 0.


 &

 &

 &
?
 	)1? The amount of time that must elapse after a refill before the Prescription can be refilled again. Required, must be
 a positive non-zero duration.


 	)

 	)+

 	).0
m
 
+.` The current status of the Prescription. Required, must not be PRESCRIPTION_STATUS_UNSPECIFIED.


 
+

 
+(

 
++-
?
 .@? When the first regimen for this Prescription begins. A regimen specifies a certain number of times that Medication
 should be taken in a certain period of time. Required.


 .

 .:

 .=?
o
 0<b The duration of a regimen for this Prescription. Required, must be a positive non-zero duration.


 0

 06

 09;
s
 2*f The number of doses of Medication that must be taken during a single regimen. Required, must be > 0.


 2

 2$

 2')
?
 5V? The points in time from the start of a regimen that Medication should be taken. Must be positive non-zero durations
 if supplied.


 5


 5#

 5$P

 5SU
?
 7.2 The date the Prescription was written. Required.


 7

 7(

 7+-
+
 ; B The status of a Prescription.



 ;
6
  =&) The Prescription status is unspecified.


  =!

  =$%
`
 ?S The Prescription is currently active (i.e. currently being taken by the Patient).


 ?

 ?
,
 A The Prescription is inactive.


 A

 Abproto3
?
!heyrenee/v1/messages/refill.protoheyrenee.v1.messagesgoogle/protobuf/timestamp.proto"?
Refill

patient_id (	R	patientId#
medication_id (	RmedicationId'
prescription_id (	RprescriptionId
	refill_id (	RrefillIdH
date_time_refilled (2.google.protobuf.TimestampRdateTimeRefilled
pharmacy_id (	R
pharmacyIdA
refill_type (2 .heyrenee.v1.messages.RefillTypeR
refillType*Q

RefillType
REFILL_TYPE_UNSPECIFIED 
REFILL_DELIVERY
REFILL_PICKUPB:Z8github.com/HeyReneeInc/proto/golang/heyrenee/v1/messagesJ?
  "

  

 

 O
	
 O
	
  )
F
 	 : A Refill represents a refill of a prescribed Medication.



 	
K
  > The ID of the Patient who the Prescription was refilled for.


  

  	

  
:
 - The ID of the Medication that was refilled.


 

 	

 
<
 / The ID of the Prescription that was refilled.


 

 	

 
$
  The ID of the Refill.


 

 	

 
:
 3- The date and time at which Refill occurred.


 

 .

 12
?
 2 The ID of the Pharmacy that supplied the Refill.


 

 	

 
"
  The type of Refill.


 

 

 
-
  "! Specifies the type of a Refill.



 
*
   Refill type is unspecified.


  

  
7
 * The Refill was delivered to the Patient.


 

 
7
 !* The Refill was picked up by the Patient.


 !

 !bproto3
?,
$heyrenee/v1/medication_service.protoheyrenee.v1*heyrenee/v1/messages/medication_dose.proto'heyrenee/v1/messages/prescription.proto!heyrenee/v1/messages/refill.protoheyrenee/v1/options/auth.proto"=
MedicationSuggestRequest!
partial_text (	RpartialText"u
MedicationSuggestResponseX
medication_suggestions (2!.heyrenee.v1.MedicationSuggestionRmedicationSuggestions"^
MedicationSuggestion#
medication_id (	RmedicationId!
display_name (	RdisplayName"c
CreatePrescriptionRequestF
prescription (2".heyrenee.v1.messages.PrescriptionRprescription"c
UpdatePrescriptionRequestF
prescription (2".heyrenee.v1.messages.PrescriptionRprescription"?
ListPrescriptionsRequest#

patient_id (	B??R	patientIdY
prescription_status (2(.heyrenee.v1.messages.PrescriptionStatusRprescriptionStatus"e
ListPrescriptionsResponseH
prescriptions (2".heyrenee.v1.messages.PrescriptionRprescriptions"K
ListMedicationDosesRequest-
prescription_id (	B??RprescriptionId"n
ListMedicationDosesResponseO
medication_doses (2$.heyrenee.v1.messages.MedicationDoseRmedicationDoses"C
ListRefillsRequest-
prescription_id (	B??RprescriptionId"M
ListRefillsResponse6
refills (2.heyrenee.v1.messages.RefillRrefills2?
MedicationServiced
MedicationSuggest%.heyrenee.v1.MedicationSuggestRequest&.heyrenee.v1.MedicationSuggestResponse" b
CreatePrescription&.heyrenee.v1.CreatePrescriptionRequest".heyrenee.v1.messages.Prescription" b
UpdatePrescription&.heyrenee.v1.UpdatePrescriptionRequest".heyrenee.v1.messages.Prescription" d
ListPrescriptions%.heyrenee.v1.ListPrescriptionsRequest&.heyrenee.v1.ListPrescriptionsResponse" j
ListMedicationDoses'.heyrenee.v1.ListMedicationDosesRequest(.heyrenee.v1.ListMedicationDosesResponse" R
ListRefills.heyrenee.v1.ListRefillsRequest .heyrenee.v1.ListRefillsResponse" B1Z/github.com/HeyReneeInc/proto/golang/heyrenee/v1J?
  j

  

 

 F
	
 F
	
  4
	
 1
	
 +
	
	 (
?
  #? MedicationService

 The MedicationService provides operations on the Medication related resources that belonging to specific Patients.

 URLs
  Scratch: heyrenee-v1-8jiz6fu6.uc.gateway.dev



 
?
  X? MedicationSuggest returns a list of MedicationSuggestions based on a partial search string. This should only be
 used for autocomplete like features and not as a full fledged Medication search method.

 TODO(mitch): Use bi-directional streaming once client streaming supported by grpc-web
 https://github.com/grpc/grpc-web/issues/24


  

  0

  ;T
M
 b@ CreatePrescription creates the provided Prescription resource.


 

 2

 =^
M
 b@ UpdatePrescription updates the provided Prescription resource.


 

 2

 =^
\
 XO ListPrescriptions lists the Prescription resources for the specified Patient.


 

 0

 ;T
\
  ^O ListMedicationDoses lists the MedicationDoses for the specified Prescription.


  

  4

  ?Z
L
 "F? ListRefills lists the Refills for the specified Prescription.


 "

 "$

 "/B
4
 & )( Request message for MedicationSuggest.



 & 
T
  (G The partial text be used to generate MedicationSuggestions. Required.


  (

  (	

  (
5
, /) Response message for MedicationSuggest.



,!
1
 .;$ The list of MedicationSuggestions.


 .


 .

 . 6

 .9:
?
3 8? A MedicationSuggestion represents a single medication suggested as an autocompletion result for the partial text in a
 MedicationSuggestRequest.



3
2
 5% The ID of the suggested Medication.


 5

 5	

 5
j
7] The name of the Medication that should be displayed to the user in the list of suggestions.


7

7	

7
5
; >) Request message for CreatePrescription.



;!
=
 =50 The Prescription resource to create. Required.


 =#

 =$0

 =34
5
A D) Request message for UpdatePrescription.



A!
=
 C50 The Prescription resource to update. Required.


 C#

 C$0

 C34
4
G L( Request message for ListPrescriptions.



G 
K
 Ib> The ID of the Patient to return Prescriptions for. Required.


 I

 I	

 I

 Ia

 ІI`
C
KB6 The status of Prescriptions that should be returned.


K)

K*=

K@A
5
O R) Response message for ListPrescriptions.



O!
)
 Q? The list of Prescriptions.


 Q


 Q,

 Q-:

 Q=>
6
U X* Request message for ListMedicationDoses.



U"
R
 WlE The ID of the Prescription to return MedicationDoses for. Required.


 W

 W	

 W

 Wk

 ІWj
7
[ ^+ Response message for ListMedicationDoses.



[#
+
 ]D The list of MedicationDoses.


 ]


 ].

 ]/?

 ]BC
.
	a d" Request message for ListRefills.



	a
J
	 cl= The ID of the Prescription to return Refills for. Required.


	 c

	 c	

	 c

	 ck

	 Іcj
/

g j# Response message for ListRefills.




g
#

 i3 The list of Refills.



 i



 i&


 i'.


 i12bproto3
?
$heyrenee/v1/messages/caregiver.protoheyrenee.v1.messagesheyrenee/v1/options/auth.proto"?
	Caregiver'
caregiver_id (	B??RcaregiverId

first_name (	R	firstName
	last_name (	RlastName#
primary_phone (	RprimaryPhone!
mobile_phone (	RmobilePhone
other_phone (	R
otherPhone
email (	Remail#
address_lines (	RaddressLines
city	 (	Rcity
state
 (	Rstate
zip (	RzipB:Z8github.com/HeyReneeInc/proto/golang/heyrenee/v1/messagesJ?
   

  

 

 O
	
 O
	
  (
?
 	  u A Caregiver represents a subuser that has access to the dashboard to monitor patients that they are caregivers for.



 	
1
  f$ The ID of the Caregiver. Required.


  

  	

  

  e

  Іd
?
 2 The legal first name of the Caregiver. Required.


 

 	

 
>
 1 The legal last name of the Caregiver. Required.


 

 	

 
O
 B The primary phone number for contacting the Caregiver. Required.


 

 	

 
D
 7 The mobile phone number for contacting the Caregiver.


 

 	

 
G
 : The alternate phone number for contacting the Caregiver.


 

 	

 
H
 ; The email address for contacting the Caregiver. Required.


 

 	

 
=
 $0 The Caregiver's permanent residential address.


 


 

 

 "#
:
 - The Caregiver's permanent residential city.


 

 	

 
;
 	. The Caregiver's permanent residential state.


 	

 		

 	
>
 
1 The Caregiver's permanent residential zip code.


 


 
	

 
bproto3
?!
,heyrenee/v1/messages/patient_caregiver.protoheyrenee.v1.messages$heyrenee/v1/messages/caregiver.protoheyrenee/v1/options/auth.proto"?
PatientCaregiver#

patient_id (	B??R	patientId#
caregiver_id (	H RcaregiverIdN
caregiver_message (2.heyrenee.v1.messages.CaregiverH RcaregiverMessage%
preferred_name (	RpreferredName`
patient_caregiver_type (2*.heyrenee.v1.messages.PatientCaregiverTypeRpatientCaregiverTypef
patient_caregiver_access (2,.heyrenee.v1.messages.PatientCaregiverAccessRpatientCaregiverAccessx
patient_caregiver_relationship (22.heyrenee.v1.messages.PatientCaregiverRelationshipRpatientCaregiverRelationshipB
	caregiver*b
PatientCaregiverType&
"PATIENT_CAREGIVER_TYPE_UNSPECIFIED "
PATIENT_CAREGIVER_TYPE_PRIMARY*?
PatientCaregiverAccess(
$PATIENT_CAREGIVER_ACCESS_UNSPECIFIED !
PATIENT_CAREGIVER_ACCESS_NONE&
"PATIENT_CAREGIVER_ACCESS_DASHBOARD*?
PatientCaregiverRelationship.
*PATIENT_CAREGIVER_RELATIONSHIP_UNSPECIFIED (
$PATIENT_CAREGIVER_RELATIONSHIP_CHILD)
%PATIENT_CAREGIVER_RELATIONSHIP_PARENT.
*PATIENT_CAREGIVER_RELATIONSHIP_GRANDPARENT-
)PATIENT_CAREGIVER_RELATIONSHIP_GRANDCHILD2
.PATIENT_CAREGIVER_RELATIONSHIP_EXTENDED_FAMILY)
%PATIENT_CAREGIVER_RELATIONSHIP_FRIEND.
*PATIENT_CAREGIVER_RELATIONSHIP_HEALTH_AIDE(
$PATIENT_CAREGIVER_RELATIONSHIP_OTHERB:Z8github.com/HeyReneeInc/proto/golang/heyrenee/v1/messagesJ?
  E

  

 

 O
	
 O
	
  .
	
 (
?
  ? A PatientCaregiver represents a relationship between a Patient and a Caregiver. The relationship indicates that the
 Caregiver is responsible in some way for managing the healthcare of the Patient.



 
P
  bC The ID of the Patient being cared for by the Caregiver. Required.


  

  	

  

  a

  І`
G
  9 The Caregiver that is caring for the Patient. Required.


  
P
 C The ID of the Caregiver that is caring for the Patient. Required.


 


 

 
t
 $g The Caregiver that is caring for the Patient. Only returned in response, must not be set in requests.


 

 

 "#
@
 3 The name that the Patient knows the Caregiver by.


 

 	

 
,
 2 The type of PatientCaregiver.


 

 -

 01
c
 6V The access that the Caregiver has to the Patient. Required, must not be unspecified.


 

 1

 45
c
 BV The relationship of the Caregiver to the Patient. Required, must not be unspecified.


 

 =

 @A
3
   %' Specifies a type of PatientCaregiver.



  
;
  "). The type of PatientCaregiver is unspecified.


  "$

  "'(
]
 $%P The Caregiver of the PatientCaregiver is the primary caregiver of the Patient.


 $ 

 $#$
s
( /g Specifies the access that a Caregiver in a PatientCaregiver has to the Patient in a PatientCaregiver.



(
+
 *+ The access is not specified.


 *&

 *)*
:
,$- The Caregiver has no access to the Patient.


,

,"#
F
.)9 The Caregiver has access to view the Patient dashboard.


.$

.'(
s
2 Eg Specifies the relationship of a Caregiver in a PatientCaregiver to the Patient in a PatientCaregiver.



2!
1
 41$ The relationship is not specified.


 4,

 4/0
4
6+' The Caregiver is the Patient's child.


6&

6)*
5
8,( The Caregiver is the Patient's parent.


8'

8*+
:
:1- The Caregiver is the Patient's grandparent.


:,

:/0
9
<0, The Caregiver is the Patient's grandchild.


<+

<./
A
>54 The Caregiver is in the Patient's extended family.


>0

>34
5
@,( The Caregiver is the Patient's friend.


@'

@*+
:
B1- The Caregiver is the Patient's health aide.


B,

B/0
E
D+8 The Caregiver has another relationship to the Patient.


D&

D)*bproto3
?
 heyrenee/v1/enums/language.protoheyrenee.v1.enums*d
Language
LANGUAGE_UNSPECIFIED 
LANGUAGE_ENGLISH
LANGUAGE_SPANISH
LANGUAGE_OTHERB7Z5github.com/HeyReneeInc/proto/golang/heyrenee/v1/enumsJ?
  

  

 

 L
	
 L
,
    Language specifies a language.



 
)
  	 Language is not specified.


  	

  	
$
  The english language.


 

 
$
  The spanish language.


 

 
?
 ? Any language not otherwise specified in this enum. Differs from LANGUAGE_UNSPECIFIED in that the language is specified
 but is not yet supported.


 

 bproto3
?6
7heyrenee/v1/messages/patient_health_questionnaire.protoheyrenee.v1.messagesheyrenee/v1/options/auth.proto heyrenee/v1/enums/language.protogoogle/protobuf/timestamp.proto"?
PatientHealthQuestionnaire#

patient_id (	B??R	patientIdE
patient_health_questionnaire_id (	RpatientHealthQuestionnaireIdR
questionnaire_language (2.heyrenee.v1.enums.LanguageRquestionnaireLanguageH
date_time_answered (2.google.protobuf.TimestampRdateTimeAnsweredp
annual_wellness_visit_status (2/.heyrenee.v1.messages.AnnualWellnessVisitStatusRannualWellnessVisitStatus`
last_annual_wellness_visit_date (2.google.protobuf.TimestampRlastAnnualWellnessVisitDateP
mammogram_status (2%.heyrenee.v1.messages.MammogramStatusRmammogramStatusJ
last_mammogram_date (2.google.protobuf.TimestampRlastMammogramDatey
colorectal_cancer_screen_status	 (22.heyrenee.v1.messages.ColorectalCancerScreenStatusRcolorectalCancerScreenStatusf
"last_colorectal_cancer_screen_date
 (2.google.protobuf.TimestampRlastColorectalCancerScreenDate|
"last_colorectal_cancer_screen_type (20.heyrenee.v1.messages.ColorectalCancerScreenTypeRlastColorectalCancerScreenTypeg
annual_flu_vaccine_status (2,.heyrenee.v1.messages.AnnualFluVaccineStatusRannualFluVaccineStatusZ
last_annual_flu_vaccine_date (2.google.protobuf.TimestampRlastAnnualFluVaccineDateo
pneumococcal_vaccine_status (2/.heyrenee.v1.messages.PneumococcalVaccineStatusRpneumococcalVaccineStatusZ
covid_vaccine_status (2(.heyrenee.v1.messages.CovidVaccineStatusRcovidVaccineStatusQ
last_covid_vaccine_date (2.google.protobuf.TimestampRlastCovidVaccineDateQ
pill_system_usage (2%.heyrenee.v1.messages.PillSystemUsageRpillSystemUsageo
daily_medication_difficulty (2/.heyrenee.v1.messages.DailyMedicationDifficultyRdailyMedicationDifficultyo
medication_copay_difficulty (2/.heyrenee.v1.messages.MedicationCopayDifficultyRmedicationCopayDifficultyu
medication_payment_difficulty (21.heyrenee.v1.messages.MedicationPaymentDifficultyRmedicationPaymentDifficulty*?
AnnualWellnessVisitStatus,
(ANNUAL_WELLNESS_VISIT_STATUS_UNSPECIFIED *
&ANNUAL_WELLNESS_VISIT_STATUS_COMPLETED.
*ANNUAL_WELLNESS_VISIT_STATUS_NOT_COMPLETED*w
MammogramStatus 
MAMMOGRAM_STATUS_UNSPECIFIED 
MAMMOGRAM_STATUS_COMPLETED"
MAMMOGRAM_STATUS_NOT_COMPLETED*?
ColorectalCancerScreenStatus/
+COLORECTAL_CANCER_SCREEN_STATUS_UNSPECIFIED -
)COLORECTAL_CANCER_SCREEN_STATUS_COMPLETED1
-COLORECTAL_CANCER_SCREEN_STATUS_NOT_COMPLETED*?
ColorectalCancerScreenType-
)COLORECTAL_CANCER_SCREEN_TYPE_UNSPECIFIED %
!COLORECTAL_CANCER_SCREEN_TYPE_FIT+
'COLORECTAL_CANCER_SCREEN_TYPE_COLOGUARD-
)COLORECTAL_CANCER_SCREEN_TYPE_COLONOSCOPY*?
AnnualFluVaccineStatus)
%ANNUAL_FLU_VACCINE_STATUS_UNSPECIFIED '
#ANNUAL_FLU_VACCINE_STATUS_COMPLETED+
'ANNUAL_FLU_VACCINE_STATUS_NOT_COMPLETED*?
PneumococcalVaccineStatus+
'PNEUMOCOCCAL_VACCINE_STATUS_UNSPECIFIED )
%PNEUMOCOCCAL_VACCINE_STATUS_COMPLETED(
$PNEUMOCOCCAL_VACCINE_STATUS_ONLY_ONE-
)PNEUMOCOCCAL_VACCINE_STATUS_NOT_COMPLETED*?
CovidVaccineStatus$
 COVID_VACCINE_STATUS_UNSPECIFIED "
COVID_VACCINE_STATUS_COMPLETED&
"COVID_VACCINE_STATUS_NOT_COMPLETED*v
PillSystemUsage#
PILL_SYSTEM_USAGE_NOT_SPECIFIED 
PILL_SYSTEM_USAGE_USES"
PILL_SYSTEM_USAGE_DOES_NOT_USE*?
DailyMedicationDifficulty-
)DAILY_MEDICATION_DIFFICULTY_NOT_SPECIFIED .
*DAILY_MEDICATION_DIFFICULTY_HAS_DIFFICULTY-
)DAILY_MEDICATION_DIFFICULTY_NO_DIFFICULTY*?
MedicationCopayDifficulty-
)MEDICATION_COPAY_DIFFICULTY_NOT_SPECIFIED .
*MEDICATION_COPAY_DIFFICULTY_HAS_DIFFICULTY-
)MEDICATION_COPAY_DIFFICULTY_NO_DIFFICULTY*?
MedicationPaymentDifficulty/
+MEDICATION_PAYMENT_DIFFICULTY_NOT_SPECIFIED ;
7MEDICATION_PAYMENT_DIFFICULTY_SKIPS_BECAUSE_CANT_AFFORD/
+MEDICATION_PAYMENT_DIFFICULTY_NO_DIFFICULTYB:Z8github.com/HeyReneeInc/proto/golang/heyrenee/v1/messagesJ?
  c

  

 

 O
	
 O
	
  (
	
 *
	
 )


 
 


 
"

  b

  

  	

  

  a

  І`

 -

 

 	(

 +,

 8

 

 3

 67

 3

 

 .

 12

 =

 

 8

 ;<

 @

 

 ;

 >?

 '

 

 "

 %&

 4

 

 /

 23

 C

 

 >

 AB

 	D

 	

 	>

 	AC

 
E

 


 
?

 
BD

 8

 

 2

 57

 >

 

 8

 ;=

 =

 

 7

 :<

 /

 

 )

 ,.

 9

 

 3

 68

 )

 

 #

 &(

 =

 

 7

 :<

 =

 

 7

 :<

 A

 

 ;

 >@


 ! %


 !

  "/

  "*

  "-.

 #-

 #(

 #+,

 $1

 $,

 $/0


' +


'

 (#

 (

 (!"

)!

)

) 

*%

* 

*#$


- 1


-!

 .2

 .-

 .01

/0

/+

/./

04

0/

023


3 8


3

 40

 4+

 4./

5(

5#

5&'

6.

6)

6,-

70

7+

7./


: >


:

 ;,

 ;'

 ;*+

<*

<%

<()

=.

=)

=,-


@ E


@

 A.

 A)

 A,-

B,

B'

B*+

C+

C&

C)*

D0

D+

D./


G K


G

 H'

 H"

 H%&

I%

I 

I#$

J)

J$

J'(


M Q


M

 N&

 N!

 N$%

O

O

O

P%

P 

P#$


S W


S

 T0

 T+

 T./

U1

U,

U/0

V0

V+

V./


	Y ]


	Y

	 Z0

	 Z+

	 Z./

	[1

	[,

	[/0

	\0

	\+

	\./



_ c



_ 


 `2


 `-


 `01


a>


a9


a<=


b2


b-


b01bproto3
?
=heyrenee/v1/messages/patient_satisfaction_questionnaire.protoheyrenee.v1.messagesheyrenee/v1/options/auth.proto heyrenee/v1/enums/language.protogoogle/protobuf/timestamp.proto"?
 PatientSatisfactionQuestionnaire#

patient_id (	B??R	patientIdQ
%patient_satisfaction_questionnaire_id (	R"patientSatisfactionQuestionnaireIdR
questionnaire_language (2.heyrenee.v1.enums.LanguageRquestionnaireLanguageH
date_time_answered (2.google.protobuf.TimestampRdateTimeAnswered#
health_status (RhealthStatus:
medical_care_satisfaction (RmedicalCareSatisfaction/
health_satisfaction (RhealthSatisfaction4
hey_renee_satisfaction (RheyReneeSatisfactionK
"primary_care_provider_satisfaction	 (RprimaryCareProviderSatisfactionB:Z8github.com/HeyReneeInc/proto/golang/heyrenee/v1/messagesJ?
  

  

 

 O
	
 O
	
  (
	
 *
	
 )
?
  | A PatientSatisfactionQuestionnaire represents the answers to a HeyRenee satisfaction questionnaire provided by a
 Patient.



 (
[
  bN The ID of the Patient who answered the satisfaction questionnaire. Required.


  

  	

  

  a

  І`
?
 3? The ID of the set of answers to the satisfaction questionnaire. Must be empty in create requests, required in update requests.


 

 	.

 12
J
 8= The language that the questionnaire was taken in. Required.


 

 3

 67
W
 3J The date and time that the Patient answered the questionnaire. Required.


 

 .

 12
^
 Q The Patient's perception of their health status on a scale of 1 to 5. Required.


 

 

 
a
 &T The Patient's satisfaction with their medical care on a scale of 1 to 5. Required.


 

 !

 $%
c
  V The Patient's satisfaction with their current health on a scale of 1 to 5. Required.


 

 

 
W
 #J The Patient's satisfaction with HeyRenee on a scale of 1 to 5. Required.


 

 

 !"
j
 /] The Patient's satisfaction with their primary care provider on a scale of 1 to 5. Required.


 

 *

 -.bproto3
?e
5heyrenee/v1/messages/patient_sdoh_questionnaire.protoheyrenee.v1.messagesheyrenee/v1/options/auth.proto heyrenee/v1/enums/language.protogoogle/protobuf/timestamp.proto"?
PatientSdohQuestionnaire#

patient_id (	B??R	patientIdA
patient_sdoh_questionnaire_id (	RpatientSdohQuestionnaireIdR
questionnaire_language (2.heyrenee.v1.enums.LanguageRquestionnaireLanguageH
date_time_answered (2.google.protobuf.TimestampRdateTimeAnsweredS
ambulatory_status (2&.heyrenee.v1.messages.AmbulatoryStatusRambulatoryStatusD
fall_history (2!.heyrenee.v1.messages.FallHistoryRfallHistoryV
walking_steadiness (2'.heyrenee.v1.messages.WalkingSteadinessRwalkingSteadinessi
walking_support_necessity (2-.heyrenee.v1.messages.WalkingSupportNecessityRwalkingSupportNecessityo
walking_assistance_advisory	 (2/.heyrenee.v1.messages.WalkingAssistanceAdvisoryRwalkingAssistanceAdvisoryr
walking_assistance_ownership
 (20.heyrenee.v1.messages.WalkingAssistanceOwnershipRwalkingAssistanceOwnershipG
falling_worry (2".heyrenee.v1.messages.FallingWorryRfallingWorry`
standing_up_assistance (2*.heyrenee.v1.messages.StandingUpAssistanceRstandingUpAssistance`
stepping_up_difficulty (2*.heyrenee.v1.messages.SteppingUpDifficultyRsteppingUpDifficulty^
feeling_in_feet_status (2).heyrenee.v1.messages.FeelingInFeetStatusRfeelingInFeetStatus?
$feeling_lightheaded_from_meds_status (26.heyrenee.v1.messages.FeelingLightheadedFromMedsStatusR feelingLightheadedFromMedsStatusJ
smoking_status (2#.heyrenee.v1.messages.SmokingStatusRsmokingStatus#
amount_smoked (	RamountSmokedW
quit_smoking_status (2'.heyrenee.v1.messages.QuitSmokingStatusRquitSmokingStatusc
alcohol_drinking_status (2+.heyrenee.v1.messages.AlcoholDrinkingStatusRalcoholDrinkingStatus0
amount_alcohol_drank (	RamountAlcoholDrankZ
quit_drinking_status (2(.heyrenee.v1.messages.QuitDrinkingStatusRquitDrinkingStatus?
*scheduling_doctors_appointments_difficulty (2=.heyrenee.v1.messages.SchedulingDoctorsAppointmentsDifficultyR'schedulingDoctorsAppointmentsDifficulty?
*getting_to_doctors_appointments_difficulty (2<.heyrenee.v1.messages.GettingToDoctorsAppointmentsDifficultyR&gettingToDoctorsAppointmentsDifficultyT
adl_help_necessity (2&.heyrenee.v1.messages.AdlHelpNecessityRadlHelpNecessity?
#getting_to_grocery_store_difficulty (25.heyrenee.v1.messages.GettingToGroceryStoreDifficultyRgettingToGroceryStoreDifficultyP
social_isolation (2%.heyrenee.v1.messages.SocialIsolationRsocialIsolationf
meal_planning_assistance (2,.heyrenee.v1.messages.MealPlanningAssistanceRmealPlanningAssistancep
weight_loss_program_interest (2/.heyrenee.v1.messages.WeightLossProgramInterestRweightLossProgramInterestA
snap_status (2 .heyrenee.v1.messages.SnapStatusR
snapStatuso
financial_assistance_status (2/.heyrenee.v1.messages.FinancialAssistanceStatusRfinancialAssistanceStatusM
food_insecurity (2$.heyrenee.v1.messages.FoodInsecurityRfoodInsecurity\
depression_frequency  (2).heyrenee.v1.messages.DepressionFrequencyRdepressionFrequency*?
AmbulatoryStatus!
AMBULATORY_STATUS_UNSPECIFIED #
AMBULATORY_STATUS_IS_AMBULATORY$
 AMBULATORY_STATUS_NOT_AMBULATORY*r
FallHistory
FALL_HISTORY_UNSPECIFIED $
 FALL_HISTORY_FALLEN_IN_LAST_YEAR
FALL_HISTORY_HAS_NOT_FALLEN*}
WalkingSteadiness"
WALKING_STEADINESS_UNSPECIFIED %
!WALKING_STEADINESS_FEELS_UNSTEADY
WALKING_STEADINESS_STEADY*?
WalkingSupportNecessity)
%WALKING_SUPPORT_NECESSITY_UNSPECIFIED ,
(WALKING_SUPPORT_NECESSITY_USES_FURNITURE"
WALKING_SUPPORT_NECESSITY_NONE*?
WalkingAssistanceAdvisory+
'WALKING_ASSISTANCE_ADVISORY_UNSPECIFIED .
*WALKING_ASSISTANCE_ADVISORY_CANE_OR_WALKER$
 WALKING_ASSISTANCE_ADVISORY_NONE*?
WalkingAssistanceOwnership,
(WALKING_ASSISTANCE_OWNERSHIP_UNSPECIFIED /
+WALKING_ASSISTANCE_OWNERSHIP_CANE_OR_WALKER%
!WALKING_ASSISTANCE_OWNERSHIP_NONE*j
FallingWorry
FALLING_WORRY_UNSPECIFIED 
FALLING_WORRY_IS_WORRIED
FALLING_WORRY_NOT_WORRIED*?
StandingUpAssistance&
"STANDING_UP_ASSISTANCE_UNSPECIFIED !
STANDING_UP_ASSISTANCE_PUSHES
STANDING_UP_ASSISTANCE_NONE*?
SteppingUpDifficulty&
"STEPPING_UP_DIFFICULTY_UNSPECIFIED )
%STEPPING_UP_DIFFICULTY_HAS_DIFFICULTY
STEPPING_UP_DIFFICULTY_NONE*?
FeelingInFeetStatus&
"FEELING_IN_FEET_STATUS_UNSPECIFIED $
 FEELING_IN_FEET_STATUS_LOST_SOME"
FEELING_IN_FEET_STATUS_HAS_ALL*?
 FeelingLightheadedFromMedsStatus4
0FEELING_LIGHTHEADED_FROM_MEDS_STATUS_UNSPECIFIED 5
1FEELING_LIGHTHEADED_FROM_MEDS_STATUS_AT_LEAST_ONE0
,FEELING_LIGHTHEADED_FROM_MEDS_STATUS_NO_MEDS*m
SmokingStatus
SMOKING_STATUS_UNSPECIFIED 
SMOKING_STATUS_SMOKES!
SMOKING_STATUS_DOES_NOT_SMOKE*?
QuitSmokingStatus#
QUIT_SMOKING_STATUS_UNSPECIFIED %
!QUIT_SMOKING_STATUS_WANTS_TO_QUIT-
)QUIT_SMOKING_STATUS_DOES_NOT_WANT_TO_QUIT*?
AlcoholDrinkingStatus'
#ALCOHOL_DRINKING_STATUS_UNSPECIFIED "
ALCOHOL_DRINKING_STATUS_DRINKS*
&ALCOHOL_DRINKING_STATUS_DOES_NOT_DRINK*?
QuitDrinkingStatus$
 QUIT_DRINKING_STATUS_UNSPECIFIED &
"QUIT_DRINKING_STATUS_WANTS_TO_QUIT.
*QUIT_DRINKING_STATUS_DOES_NOT_WANT_TO_QUIT*?
'SchedulingDoctorsAppointmentsDifficulty:
6SCHEDULING_DOCTORS_APPOINTMENTS_DIFFICULTY_UNSPECIFIED =
9SCHEDULING_DOCTORS_APPOINTMENTS_DIFFICULTY_HAS_DIFFICULTY<
8SCHEDULING_DOCTORS_APPOINTMENTS_DIFFICULTY_NO_DIFFICULTY*?
&GettingToDoctorsAppointmentsDifficulty:
6GETTING_TO_DOCTORS_APPOINTMENTS_DIFFICULTY_UNSPECIFIED C
?GETTING_TO_DOCTORS_APPOINTMENTS_DIFFICULTY_TROUBLE_FINDING_RIDE<
8GETTING_TO_DOCTORS_APPOINTMENTS_DIFFICULTY_NO_DIFFICULTY*x
AdlHelpNecessity"
ADL_HELP_NECESSITY_UNSPECIFIED 
ADL_HELP_NECESSITY_NEEDED!
ADL_HELP_NECESSITY_NOT_NEEDED*?
GettingToGroceryStoreDifficulty3
/GETTING_TO_GROCERY_STORE_DIFFICULTY_UNSPECIFIED 4
0GETTING_TO_GROCERY_STORE_DIFFICULTY_UNABLE_ALONE2
.GETTING_TO_GROCERY_STORE_DIFFICULTY_ABLE_ALONE*?
SocialIsolation 
SOCIAL_ISOLATION_UNSPECIFIED #
SOCIAL_ISOLATION_FEELS_ISOLATED+
'SOCIAL_ISOLATION_DOES_NOT_FEEL_ISOLATED*?
MealPlanningAssistance(
$MEAL_PLANNING_ASSISTANCE_UNSPECIFIED "
MEAL_PLANNING_ASSISTANCE_WANTS*
&MEAL_PLANNING_ASSISTANCE_DOES_NOT_WANT*?
WeightLossProgramInterest,
(WEIGHT_LOSS_PROGRAM_INTEREST_UNSPECIFIED +
'WEIGHT_LOSS_PROGRAM_INTEREST_INTERESTED/
+WEIGHT_LOSS_PROGRAM_INTEREST_NOT_INTERESTED*_

SnapStatus
SNAP_STATUS_UNSPECIFIED 
SNAP_STATUS_ON_SNAP
SNAP_STATUS_NOT_ON_SNAP*?
FinancialAssistanceStatus+
'FINANCIAL_ASSISTANCE_STATUS_UNSPECIFIED =
9FINANCIAL_ASSISTANCE_STATUS_RECEIVES_FINANCIAL_ASSISTANCE7
3FINANCIAL_ASSISTANCE_STATUS_NO_FINANCIAL_ASSISTANCE*k
FoodInsecurity
FOOD_INSECURITY_UNSPECIFIED 
FOOD_INSECURITY_INSECURE
FOOD_INSECURITY_SECURE*?
DepressionFrequency$
 DEPRESSION_FREQUENCY_UNSPECIFIED !
DEPRESSION_FREQUENCY_FREQUENT#
DEPRESSION_FREQUENCY_INFREQUENTB:Z8github.com/HeyReneeInc/proto/golang/heyrenee/v1/messagesJ?-
  ?

  

 

 O
	
 O
	
  (
	
 *
	
 )


 
 +


 
 

  b

  

  	

  

  a

  І`

 +

 

 	&

 )*

 8

 

 3

 67

 3

 

 .

 12

 )

 

 $

 '(

 

 

 

 

 +

 

 &

 )*

 8

 

 3

 67

 <

 

 7

 :;

 	?

 	

 	9

 	<>

 
"

 


 


 
!

 3

 

 -

 02

 3

 

 -

 02

 2

 

 ,

 /1

 M

 "

 #G

 JL

 $

 

 

 !#

 

 

 	

 

 -

 

 '

 *,

 5

 

 /

 24

 #

 

 	

  "

 /

 

 )

 ,.

  Z

  )

  *T

  WY

 !Y

 !(

 !)S

 !VX

 "+

 "

 "%

 "(*

 #K

 #!

 #"E

 #HJ

 $(

 $

 $"

 $%'

 %7

 %

 %1

 %46

 &>

 &

 &8

 &;=

 '

 '

 '

 '

 (=

 (

 (7

 (:<

 )&

 )

 ) 

 )#%

 *0

 *

 **

 *-/


 - 1


 -

  .$

  .

  ."#

 /&

 /!

 /$%

 0'

 0"

 0%&


3 7


3

 4

 4

 4

5'

5"

5%&

6"

6

6 !


9 =


9

 :%

 : 

 :#$

;(

;#

;&'

< 

<

<


? C


?

 @,

 @'

 @*+

A/

A*

A-.

B%

B 

B#$


E I


E

 F.

 F)

 F,-

G1

G,

G/0

H'

H"

H%&


K O


K

 L/

 L*

 L-.

M2

M-

M01

N(

N#

N&'


Q U


Q

 R 

 R

 R

S

S

S

T 

T

T


W [


W

 X)

 X$

 X'(

Y$

Y

Y"#

Z"

Z

Z !


] a


]

 ^)

 ^$

 ^'(

_,

_'

_*+

`"

`

` !


	c g


	c

	 d)

	 d$

	 d'(

	e'

	e"

	e%&

	f%

	f 

	f#$



i m



i%


 j7


 j2


 j56


k8


k3


k67


l3


l.


l12


o s


o

 p!

 p

 p 

q

q

q

r$

r

r"#


u y


u

 v&

 v!

 v$%

w(

w#

w&'

x0

x+

x./


{ 


{

 |*

 |%

 |()

}%

} 

}#$

~-

~(

~+,

? ?

?

 ?'

 ?"

 ?%&

?)

?$

?'(

?1

?,

?/0

? ?

?,

 ?=

 ?8

 ?;<

?@

?;

?>?

??

?:

?=>

? ?

?+

 ?=

 ?8

 ?;<

?F

?A

?DE

??

?:

?=>

? ?

?

 ?%

 ? 

 ?#$

? 

?

?

?$

?

?"#

? ?

?$

 ?6

 ?1

 ?45

?7

?2

?56

?5

?0

?34

? ?

?

 ?#

 ?

 ?!"

?&

?!

?$%

?.

?)

?,-

? ?

?

 ?+

 ?&

 ?)*

?%

? 

?#$

?-

?(

?+,

? ?

?

 ?/

 ?*

 ?-.

?.

?)

?,-

?2

?-

?01

? ?

?

 ?

 ?

 ?

?

?

?

?

?

?

? ?

?

 ?.

 ?)

 ?,-

?@

?;

?>?

?:

?5

?89

? ?

?

 ?"

 ?

 ? !

?

?

?

?

?

?

? ?

?

 ?'

 ?"

 ?%&

?$

?

?"#

?&

?!

?$%bproto3
?
-heyrenee/v1/messages/patient_assessment.protoheyrenee.v1.messagesheyrenee/v1/options/auth.proto heyrenee/v1/enums/language.protogoogle/protobuf/timestamp.proto"?
PatientAssessment#

patient_id (	B??R	patientId2
patient_assessment_id (	RpatientAssessmentIdM
assessment_type (2$.heyrenee.v1.messages.AssessmentTypeRassessmentTypeL
assessment_language (2.heyrenee.v1.enums.LanguageRassessmentLanguageB
date_time_taken (2.google.protobuf.TimestampRdateTimeTaken
score (Rscore*
AssessmentType
ASSESSMENT_TYPE_UNSPECIFIED #
ASSESSMENT_TYPE_HEALTH_LITERACY'
#ASSESSMENT_TYPE_TECHNOLOGY_LITERACYB:Z8github.com/HeyReneeInc/proto/golang/heyrenee/v1/messagesJ?	
  "

  

 

 O
	
 O
	
  (
	
 *
	
 )
]
  Q A PatientAssessment represents the results of an assessment taken by a Patient.



 
H
  b; The ID of the Patient that took the assessment. Required.


  

  	

  

  a

  І`
n
 #a The ID of the PatientAssessment. Must be empty in create requests, required in update requests.


 

 	

 !"
?
 %2 The type of assessment that was taken. Required.


 

  

 #$
G
 5: The language that the assessment was taken in. Required.


 

 0

 34
I
 0< The date and time that the assessment was taken. Required.


 

 +

 ./
E
 8 The score that the Patient received on the assessment.


 

 	

 
-
  "! Specifies a type of assessment.



 
2
  "% The assessment type is unspecified.


  

   !
>
 &1 The assessment is a health literacy assessment.


 !

 $%
B
 !*5 The assessment is a technology literacy assessment.


 !%

 !()bproto3
?G
!heyrenee/v1/patient_service.protoheyrenee.v1,heyrenee/v1/messages/patient_caregiver.proto7heyrenee/v1/messages/patient_health_questionnaire.proto=heyrenee/v1/messages/patient_satisfaction_questionnaire.proto5heyrenee/v1/messages/patient_sdoh_questionnaire.proto-heyrenee/v1/messages/patient_assessment.protoheyrenee/v1/options/auth.proto"t
CreatePatientCaregiverRequestS
patient_caregiver (2&.heyrenee.v1.messages.PatientCaregiverRpatientCaregiver"t
UpdatePatientCaregiverRequestS
patient_caregiver (2&.heyrenee.v1.messages.PatientCaregiverRpatientCaregiver"C
ListPatientCaregiversRequest#

patient_id (	B??R	patientId"v
ListPatientCaregiversResponseU
patient_caregivers (2&.heyrenee.v1.messages.PatientCaregiverRpatientCaregivers"x
CreatePatientAssessmentRequestV
patient_assessment (2'.heyrenee.v1.messages.PatientAssessmentRpatientAssessment"D
ListPatientAssessmentsRequest#

patient_id (	B??R	patientId"z
ListPatientAssessmentsResponseX
patient_assessments (2'.heyrenee.v1.messages.PatientAssessmentRpatientAssessments"?
-CreatePatientSatisfactionQuestionnaireRequest?
"patient_satisfaction_questionnaire (26.heyrenee.v1.messages.PatientSatisfactionQuestionnaireR patientSatisfactionQuestionnaire"S
,ListPatientSatisfactionQuestionnairesRequest#

patient_id (	B??R	patientId"?
-ListPatientSatisfactionQuestionnairesResponse?
#patient_satisfaction_questionnaires (26.heyrenee.v1.messages.PatientSatisfactionQuestionnaireR!patientSatisfactionQuestionnaires"?
'CreatePatientHealthQuestionnaireRequestr
patient_health_questionnaire (20.heyrenee.v1.messages.PatientHealthQuestionnaireRpatientHealthQuestionnaire"M
&ListPatientHealthQuestionnairesRequest#

patient_id (	B??R	patientId"?
'ListPatientHealthQuestionnairesResponset
patient_health_questionnaires (20.heyrenee.v1.messages.PatientHealthQuestionnaireRpatientHealthQuestionnaires"?
%CreatePatientSdohQuestionnaireRequestl
patient_sdoh_questionnaire (2..heyrenee.v1.messages.PatientSdohQuestionnaireRpatientSdohQuestionnaire"K
$ListPatientSdohQuestionnairesRequest#

patient_id (	B??R	patientId"?
%ListPatientSdohQuestionnairesResponsen
patient_sdoh_questionnaires (2..heyrenee.v1.messages.PatientSdohQuestionnaireRpatientSdohQuestionnaires2?
PatientServicen
CreatePatientCaregiver*.heyrenee.v1.CreatePatientCaregiverRequest&.heyrenee.v1.messages.PatientCaregiver" n
UpdatePatientCaregiver*.heyrenee.v1.UpdatePatientCaregiverRequest&.heyrenee.v1.messages.PatientCaregiver" p
ListPatientCaregivers).heyrenee.v1.ListPatientCaregiversRequest*.heyrenee.v1.ListPatientCaregiversResponse" q
CreatePatientAssessment+.heyrenee.v1.CreatePatientAssessmentRequest'.heyrenee.v1.messages.PatientAssessment" s
ListPatientAssessments*.heyrenee.v1.ListPatientAssessmentsRequest+.heyrenee.v1.ListPatientAssessmentsResponse" ?
&CreatePatientSatisfactionQuestionnaire:.heyrenee.v1.CreatePatientSatisfactionQuestionnaireRequest6.heyrenee.v1.messages.PatientSatisfactionQuestionnaire" ?
%ListPatientSatisfactionQuestionnaires9.heyrenee.v1.ListPatientSatisfactionQuestionnairesRequest:.heyrenee.v1.ListPatientSatisfactionQuestionnairesResponse" ?
 CreatePatientHealthQuestionnaire4.heyrenee.v1.CreatePatientHealthQuestionnaireRequest0.heyrenee.v1.messages.PatientHealthQuestionnaire" ?
ListPatientHealthQuestionnaires3.heyrenee.v1.ListPatientHealthQuestionnairesRequest4.heyrenee.v1.ListPatientHealthQuestionnairesResponse" ?
CreatePatientSdohQuestionnaire2.heyrenee.v1.CreatePatientSdohQuestionnaireRequest..heyrenee.v1.messages.PatientSdohQuestionnaire" ?
ListPatientSdohQuestionnaires1.heyrenee.v1.ListPatientSdohQuestionnairesRequest2.heyrenee.v1.ListPatientSdohQuestionnairesResponse" B1Z/github.com/HeyReneeInc/proto/golang/heyrenee/v1J?)
  ?

  

 

 F
	
 F
	
  6
	
 A
	
 G
	
	 ?
	

 7
	
 (
?
  0? PatientService

 The PatientService provides functionality for managing Patient subusers. Specifically, it provides methods for
 storing and accessing the results of Patient assessments, storing and accessing the results of Patient
 questionnaires, and managing relationships between Patients and other subusers.

 URLs
  Scratch: heyrenee-v1-8jiz6fu6.uc.gateway.dev



 
K
  n> CreatePatientCaregiver creates a specified PatientCaregiver.


  

  :

  Ej
K
 n> UpdatePatientCaregiver updates a specified PatientCaregiver.


 

 :

 Ej
T
 dG ListPatientCaregivers lists PatientCaregivers for a specific Patient.


 

 8

 C`
M
 q@ CreatePatientAssessment creates a specified PatientAssessment.


 

 <

 Gm
V
  gI ListPatientAssessments lists PatientAssessments for a specific Patient.


  

  :

  Ec
l
 #?^ CreatePatientSatisfactionQuestionnaire creates a specified PatientSatisfactionQuestionnaire.


 #,

 #-Z

 #e?
u
 %?g ListPatientSatisfactionQuestionnaires lists PatientSatisfactionQuestionnaires for a specific Patient.


 %+

 %,X

 %c?
`
 (?R CreatePatientHealthQuestionnaire creates a specified PatientHealthQuestionnaire.


 (&

 ('N

 (Y?
i
 *?[ ListPatientHealthQuestionnaires lists PatientHealthQuestionnaires for a specific Patient.


 *%

 *&L

 *W~
\
 	-?N CreatePatientSdohQuestionnaire creates a specified PatientSdohQuestionnaire.


 	-$

 	-%J

 	-U?
d
 
/|W ListPatientSdohQuestionnaires lists PatientSdohQuestionnaires for a specific Patient.


 
/#

 
/$H

 
/Sx
9
 3 6- Request message for CreatePatientCaregiver.



 3%
8
  5>+ The PatientCaregiver to create. Required.


  5'

  5(9

  5<=
9
9 <- Request message for UpdatePatientCaregiver.



9%
8
 ;>+ The PatientCaregiver to update. Required.


 ;'

 ;(9

 ;<=
8
? B, Request message for ListPatientCaregivers.



?$
O
 AbB The ID of the Patient to return PatientCaregivers for. Required.


 A

 A	

 A

 Aa

 ІA`
9
E H- Response message for ListPatientCaregivers.



E%
-
 GH  The list of PatientCaregivers.


 G


 G0

 G1C

 GFG
:
K N. Request message for CreatePatientAssessment.



K&
9
 M@, The PatientAssessment to create. Required.


 M(

 M);

 M>?
9
Q T- Request message for ListPatientAssessments.



Q%
P
 SbC The ID of the Patient to return PatientAssessments for. Required.


 S

 S	

 S

 Sa

 ІS`
:
W Z. Response message for ListPatientAssessments.



W&
.
 YJ! The list of PatientAssessments.


 Y


 Y1

 Y2E

 YHI
I
] `= Request message for CreatePatientSatisfactionQuestionnaire.



]5
H
 __; The PatientSatisfactionQuestionnaire to create. Required.


 _7

 _8Z

 _]^
H
c f< Request message for ListPatientSatisfactionQuestionnaires.



c4
_
 ebR The ID of the Patient to return PatientSatisfactionQuestionnaires for. Required.


 e

 e	

 e

 ea

 Іe`
I
	i l= Response message for ListPatientSatisfactionQuestionnaires.



	i5
=
	 ki0 The list of PatientSatisfactionQuestionnaires.


	 k


	 k@

	 kAd

	 kgh
C

o r7 Request message for CreatePatientHealthQuestionnaire.




o/
B

 qS5 The PatientHealthQuestionnaire to create. Required.



 q1


 q2N


 qQR
B
u x6 Request message for ListPatientHealthQuestionnaires.



u.
Y
 wbL The ID of the Patient to return PatientHealthQuestionnaires for. Required.


 w

 w	

 w

 wa

 Іw`
C
{ ~7 Response message for ListPatientHealthQuestionnaires.



{/
7
 }]* The list of PatientHealthQuestionnaires.


 }


 }:

 };X

 }[\
C
? ?5 Request message for CreatePatientSdohQuestionnaire.


?-
A
 ?O3 The PatientSdohQuestionnaire to create. Required.


 ?/

 ?0J

 ?MN
B
? ?4 Request message for ListPatientSdohQuestionnaires.


?,
X
 ?bJ The ID of the Patient to return PatientSdohQuestionnaires for. Required.


 ?

 ?	

 ?

 ?a

 І?`
C
? ?5 Response message for ListPatientSdohQuestionnaires.


?-
6
 ?Y( The list of PatientSdohQuestionnaires.


 ?


 ?8

 ?9T

 ?WXbproto3
?
+heyrenee/v1/messages/patient_provider.protoheyrenee.v1.messages#heyrenee/v1/messages/provider.proto!heyrenee/v1/enums/specialty.protoheyrenee/v1/options/auth.proto"?
PatientProvider#

patient_id (	B??R	patientId!
provider_id (	H R
providerIdK
provider_message (2.heyrenee.v1.messages.ProviderH RproviderMessage

cell_phone (	R	cellPhone1
contact_instructions (	RcontactInstructions!
contact_info (	RcontactInfoc
patient_provider_status (2+.heyrenee.v1.messages.PatientProviderStatusRpatientProviderStatus]
patient_provider_type (2).heyrenee.v1.messages.PatientProviderTypeRpatientProviderType:
	specialty	 (2.heyrenee.v1.enums.SpecialtyR	specialtyB

provider*|
PatientProviderStatus'
#PATIENT_PROVIDER_STATUS_UNSPECIFIED 
PATIENT_PROVIDER_ACTIVE
PATIENT_PROVIDER_INACTIVE*_
PatientProviderType%
!PATIENT_PROVIDER_TYPE_UNSPECIFIED !
PATIENT_PROVIDER_TYPE_PRIMARYB:Z8github.com/HeyReneeInc/proto/golang/heyrenee/v1/messagesJ?
  6

  

 

 O
	
 O
	
  -
	
 +
	
 (
?
  #? A PatientProvider represents a relationship between a Patient and Provider resource. The relationship indicates that
 the Provider is either currently or has in the past provided healthcare services for the Patient.



 
g
  bZ The ID of the Patient that is receiving healthcare services from the Provider. Required.


  

  	

  

  a

  І`
\
  N The Provider that is providing healthcare services to the Patient. Required.


  
e
 X The ID of the Provider that is providing healthcare services to the Patient. Required.


 


 

 
?
 "~ The Provider that is providing healthcare services to the Patient. Only returned in responses, must not be set
 in requests.


 

 

  !
c
 V The cell phone number specifically for this Patient to use to contact this Provider.


 

 	

 
c
 "V Contact instructions specifically for how this Patient should contact this Provider.


 

 	

  !
_
 R Contact information specifically for this Patient to contact this Provider with.


 

 	

 
l
 4_ The status of the PatientProvider. Required, must not be PATIENT_PROVIDER_STATUS_UNSPECIFIED.


 

 /

 23
+
  0 The type of PatientProvider.


  

  +

  ./
Q
 ",D The specialty that this Provider specifically offers this Patient.


 "

 "'

 "*+
3
 & .' Specifies a PatientProvider's status.



 &
;
  (*. The PatientProvider's status is unspecified.


  (%

  (()
?
 *v The Provider of the PatientProvider is actively providing healthcare services to the Patient of the PatientProvider.


 *

 *
?
 - ? The Provider of the PatientProvider is not actively providing healthcare services to the Patient of the
 PatientProvider, but has at some point in the past.


 -

 -
1
1 6% Specifies a PatientProvider's type.



1
9
 3(, The PatientProvider's type in unspecified.


 3#

 3&'
r
5$e The Provider of the PatientProvider is the primary provider for the Patient of the PatientProvider.


5

5"#bproto3
?
"heyrenee/v1/provider_service.protoheyrenee.v1+heyrenee/v1/messages/patient_provider.proto#heyrenee/v1/messages/provider.proto!heyrenee/v1/enums/specialty.protoheyrenee/v1/options/auth.proto"?
SearchProvidersRequest

first_name (	R	firstName
	last_name (	RlastName:
	specialty (2.heyrenee.v1.enums.SpecialtyR	specialty
city (	Rcity
state (	Rstate"W
SearchProvidersResponse<
	providers (2.heyrenee.v1.messages.ProviderR	providers"p
CreatePatientProviderRequestP
patient_provider (2%.heyrenee.v1.messages.PatientProviderRpatientProvider"p
UpdatePatientProviderRequestP
patient_provider (2%.heyrenee.v1.messages.PatientProviderRpatientProvider"?
ListPatientProvidersRequest#

patient_id (	B??R	patientIdc
patient_provider_status (2+.heyrenee.v1.messages.PatientProviderStatusRpatientProviderStatus]
patient_provider_type (2).heyrenee.v1.messages.PatientProviderTypeRpatientProviderType"r
ListPatientProvidersResponseR
patient_providers (2%.heyrenee.v1.messages.PatientProviderRpatientProviders2?
ProviderService^
SearchProviders#.heyrenee.v1.SearchProvidersRequest$.heyrenee.v1.SearchProvidersResponse" k
CreatePatientProvider).heyrenee.v1.CreatePatientProviderRequest%.heyrenee.v1.messages.PatientProvider" k
UpdatePatientProvider).heyrenee.v1.UpdatePatientProviderRequest%.heyrenee.v1.messages.PatientProvider" m
ListPatientProviders(.heyrenee.v1.ListPatientProvidersRequest).heyrenee.v1.ListPatientProvidersResponse" B1Z/github.com/HeyReneeInc/proto/golang/heyrenee/v1J?
  G

  

 

 F
	
 F
	
  5
	
 -
	
 +
	
	 (
?
  ? ProviderService

 The Provider service provides operations on Provider resources as well as PatientProviders which represent a
 relationship between a Patient user and a Provider resource.

 URLs
  Scratch: heyrenee-v1-8jiz6fu6.uc.gateway.dev



 
c
  RV SearchProviders returns a list of Providers based on the provided search parameters.


  

  ,

  7N
I
 k< CreatePatientProvider creates a specified PatientProvider.


 

 8

 Cg
I
 k< UpdatePatientProvider updates a specified PatientProvider.


 

 8

 Cg
S
 aF ListPatientProviders lists PatientProviders for a specified Patient.


 

 6

 A]
2
  %& Request message for SearchProviders.



 

   

   

   	

   

 !

 !

 !	

 !

 ",

 "

 "'

 "*+

 #

 #

 #	

 #

 $

 $

 $	

 $
3
( +' Response message for SearchProviders.



(
%
 *7 The list of Providers.


 *


 *(

 *)2

 *56
8
. 1, Request message for CreatePatientProvider.



.$
7
 0<* The PatientProvider to create. Required.


 0&

 0'7

 0:;
8
4 7, Request message for UpdatePatientProvider.



4$
7
 6<* The PatientProvider to update. Required.


 6&

 6'7

 6:;
7
: A+ Request message for ListPatientProviders.



:#
N
 <bA The ID of the Patient to return PatientProviders for. Required.


 <

 <	

 <

 <a

 І<`
q
>Id The PatientProvider status. If specified, only PatientProviders with this status will be returned.


>,

>-D

>GH
k
@E^ The PatientProvider type. If specified, only PatientProviders of this type will be returned.


@*

@+@

@CD
8
D G, Response message for ListPatientProviders.



D$
,
 FF The list of PatientProviders.


 F


 F/

 F0A

 FDEbproto3
? 
'heyrenee/v1/messages/rpm_schedule.protoheyrenee.v1.messagesgoogle/protobuf/duration.protogoogle/protobuf/timestamp.proto#heyrenee/v1/messages/provider.proto"?
RpmSchedule

patient_id (	R	patientId&
rpm_schedule_id (	RrpmScheduleIdQ
rpm_schedule_type (2%.heyrenee.v1.messages.RpmScheduleTypeRrpmScheduleTypea
first_measurement_regimen_start (2.google.protobuf.TimestampRfirstMeasurementRegimenStart[
measurement_regimen_duration (2.google.protobuf.DurationRmeasurementRegimenDuration8
measurements_per_regimen (RmeasurementsPerRegimenq
(measurement_durations_from_regimen_start (2.google.protobuf.DurationR$measurementDurationsFromRegimenStart!
provider_id (	H R
providerIdK
provider_message	 (2.heyrenee.v1.messages.ProviderH RproviderMessageW
rpm_schedule_status
 (2'.heyrenee.v1.messages.RpmScheduleStatusRrpmScheduleStatusB

provider*?
RpmScheduleType!
RPM_SCHEDULE_TYPE_UNSPECIFIED  
RPM_SCHEDULE_TYPE_HEART_RATE$
 RPM_SCHEDULE_TYPE_BLOOD_PRESSURE
RPM_SCHEDULE_TYPE_PULSE
RPM_SCHEDULE_TYPE_SP_O2
RPM_SCHEDULE_TYPE_WEIGHT
RPM_SCHEDULE_TYPE_GLUCOSE*l
RpmScheduleStatus#
RPM_SCHEDULE_STATUS_UNSPECIFIED 
RPM_SCHEDULE_ACTIVE
RPM_SCHEDULE_INACTIVEB:Z8github.com/HeyReneeInc/proto/golang/heyrenee/v1/messagesJ?
  C

  

 

 O
	
 O
	
  (
	
 )
	
 -
?
  '? An RpmSchedule represents a schedule of when a patient should use remote patient monitoring (RPM) devices to take
 specific types of RPM measurements. Each schedule specifies a certain number of measurements that should be taken
 within a certain period of time and at certain times (e.g. two measurements in 24 hours at 8 hours and 16 hours).
 A regimen refers to single period where measurements are taken (e.g. the two measurements taken on 10/24 at 8:00 and
 16:00).



 
?
  2 The ID of the Patient that this schedule is for.


  

  	

  
*
  The ID of the RpmSchedules.


 

 	

 
'
 ( The type of RpmSchedule.


 

 #

 &'
D
 @7 The start time of the first regimen of this schedule.


 

 ;

 >?
>
 <1 The duration of a regimen within this schedule.


 

 7

 :;
U
 %H The number of measurements that must be taken during a single regimen.


 

  

 #$
f
 QY The lengths of time after the beginning of a regimen that measurements should be taken.


 


 #

 $L

 OP
;
  $- The Provider who established this schedule.


  
D
 !7 The ID of the Provider who established this schedule.


 !


 !

 !
:
 #"- The Provider who established this schedule.


 #

 #

 # !
-
 	&-  The status of the RpmSchedule.


 	&

 	&'

 	&*,
C
 * 97 RpmScheduleType specifies the type of an RpmSchedule.



 *
5
  ,$( The RpmSchedule type is not specified.


  ,

  ,"#
<
 .#/ The RpmSchedule is for heart rate monitoring.


 .

 .!"
@
 0'3 The RpmSchedule is for blood pressure monitoring.


 0"

 0%&
7
 2* The RpmSchedule is for pulse monitoring.


 2

 2
6
 4) The RpmSchedule is for Sp02 monitoring.


 4

 4
8
 6+ The RpmSchedule is for weight monitoring.


 6

 6
9
 8 , The RpmSchedule is for glucose monitoring.


 8

 8
G
< C; RpmScheduleStatus specifies the status of an RpmSchedule/



<
7
 >&* The RpmSchedule status is not specified.


 >!

 >$%
}
@p The RpmSchedule is currently active (i.e. the Patient is currently taking RPM measurements for this schedule).


@

@
+
B The RpmSchedule is inactive.


B

Bbproto3
?
*heyrenee/v1/messages/rpm_measurement.protoheyrenee.v1.messagesgoogle/protobuf/timestamp.proto"?
RpmMeasurement

patient_id (	R	patientId&
rpm_schedule_id (	RrpmScheduleId,
rpm_measurement_id (	RrpmMeasurementIdb
rpm_measurement_results (2*.heyrenee.v1.messages.RpmMeasurementResultRrpmMeasurementResults9

time_taken (2.google.protobuf.TimestampR	timeTaken"?
RpmMeasurementResult9
rpm_measurement_result_id (	RrpmMeasurementResultId
value (Rvaluem
rpm_measurement_result_unit (2..heyrenee.v1.messages.RpmMeasurementResultUnitRrpmMeasurementResultUnitm
rpm_measurement_result_type (2..heyrenee.v1.messages.RpmMeasurementResultTypeRrpmMeasurementResultType*?
RpmMeasurementResultUnit+
'RPM_MEASUREMENT_RESULT_UNIT_UNSPECIFIED #
RPM_MEASUREMENT_RESULT_UNIT_BPM#
RPM_MEASUREMENT_RESULT_UNIT_LBS*
&RPM_MEASUREMENT_RESULT_UNIT_PERCENTAGE%
!RPM_MEASUREMENT_RESULT_UNIT_MG_DL%
!RPM_MEASUREMENT_RESULT_UNIT_MM_HG'
#RPM_MEASUREMENT_RESULT_UNIT_CELSIUS*?
RpmMeasurementResultType+
'RPM_MEASUREMENT_RESULT_TYPE_UNSPECIFIED &
"RPM_MEASUREMENT_RESULT_TYPE_WEIGHT$
 RPM_MEASUREMENT_RESULT_TYPE_SPO2#
RPM_MEASUREMENT_RESULT_TYPE_BPM'
#RPM_MEASUREMENT_RESULT_TYPE_GLUCOSE#
RPM_MEASUREMENT_RESULT_TYPE_DIA#
RPM_MEASUREMENT_RESULT_TYPE_SYS$
 RPM_MEASUREMENT_RESULT_TYPE_TEMPB:Z8github.com/HeyReneeInc/proto/golang/heyrenee/v1/messagesJ?
  F

  

 

 O
	
 O
	
  )
m
 	 a An RpmMeasurement represents a single event where an RPM device was used to take a measurement.



 	
B
  5 The ID of the Patient that this measurement is for.


  

  	

  
J
 = The ID of the RpmSchedule that this measurement is part of.


 

 	

 
,
   The ID of the RpmMeasurement.


 

 	

 
<
 </ A list of the results of this RpmMeasurement.


 


 

  7

 :;
;
 +. The time at which the measurement was taken.


 

 &

 )*
d
  X An RpmMeasurementResult represents a single value output from an RpmMeasurement event.




2
 '% The ID of the RpmMeasurementResult.


 

 	"

 %&
'
 The value of the result.




	


;
;. The unit that the value of the result is in.




6

9:
0
;# The type of RpmMeasurementResult.




6

9:
D
 # 28 Specifies the unit that an RpmMeasurementResult is in.



 #
)
  %. The unit is not specified.


  %)

  %,-
2
 '&% The unit is beats per minute (BPM).


 '!

 '$%
(
 )& The unit is pounds (lbs).


 )!

 )$%
(
 +- The unit is a percentage.


 +(

 ++,
;
 -(. The unit is milligram per decilitre (mg/dL).


 -#

 -&'
9
 /(, The unit is millimeters of mercury (mmHg).


 /#

 /&'
#
 1* The unit is celsius.


 1%

 1()
<
5 F0 Specifies the type of an RpmMeasurementResult.



5
)
 7. The type is not specified.


 7)

 7,-
2
9)% The result is a weight measurement.


9$

9'(
8
;'+ The result is a blood oxygen measurement.


;"

;%&
6
=&) The result is a heart rate measurement.


=!

=$%
7
?** The result is a blood sugar measurement.


?%

?()
D
A&7 The result is a diastolic blood pressure measurement.


A!

A$%
C
C&6 The result is a systolic blood pressure measurement.


C!

C$%
<
E'/ The result is a body temperature measurement.


E"

E%&bproto3
?
heyrenee/v1/rpm_service.protoheyrenee.v1'heyrenee/v1/messages/rpm_schedule.proto*heyrenee/v1/messages/rpm_measurement.protoheyrenee/v1/options/auth.proto"?
ListRpmSchedulesRequest#

patient_id (	B??R	patientIdW
rpm_schedule_status (2'.heyrenee.v1.messages.RpmScheduleStatusRrpmScheduleStatus"b
ListRpmSchedulesResponseF
rpm_schedules (2!.heyrenee.v1.messages.RpmScheduleRrpmSchedules"J
ListRpmMeasurementsRequest,
rpm_schedule_id (	B??RrpmScheduleId"n
ListRpmMeasurementsResponseO
rpm_measurements (2$.heyrenee.v1.messages.RpmMeasurementRrpmMeasurements2?

RpmServicea
ListRpmSchedules$.heyrenee.v1.ListRpmSchedulesRequest%.heyrenee.v1.ListRpmSchedulesResponse" j
ListRpmMeasurements'.heyrenee.v1.ListRpmMeasurementsRequest(.heyrenee.v1.ListRpmMeasurementsResponse" B1Z/github.com/HeyReneeInc/proto/golang/heyrenee/v1J?
  2

  

 

 F
	
 F
	
  1
	
 4
	
 (
?
  ? RpmService

 The RpmService provides access to resources related to Remote Patient Monitoring (RPM).

 URLs
  Scratch: heyrenee-v1-8jiz6fu6.uc.gateway.dev



 
?
  U? ListRpmSchedules retrieves a list of a Patient's RpmSchedules. These schedules specify when a patient should use
 RPM devices to take different RPM measurements.


  

  .

  9Q
?
 ^? ListRpmMeasurements retrieves a list of RpmMeasurements for a specific RpmSchedule. These measurements represent a
 single event where an RPM device was used to take a measurement. The RpmMeasurement contains a list of RpmResults
 where each individual result of the measurement event is listed.


 

 4

 ?Z
3
   ' Request message for ListRpmSchedules.



 
J
  b= The ID of the Patient to return RpmSchedules for. Required.


  

  	

  

  a

  І`
k
 A^ The status of the RpmSchedules. If specified, only RpmSchedules of this status are returned.


 (

 )<

 ?@
4
# &( Response message for ListRpmSchedules.



# 
(
 %> The list of RpmSchedules.


 %


 %+

 %,9

 %<=
6
) ,* Request message for ListRpmMeasurements.



)"
Q
 +lD The ID of the RpmSchedule to return RpmMeasurements for. Required.


 +

 +	

 +

 +k

 І+j
7
/ 2+ Response message for ListRpmMeasurements.



/#
+
 1D The list of RpmMeasurements.


 1


 1.

 1/?

 1BCbproto3
?$
"heyrenee/v1/messages/patient.protoheyrenee.v1.messagesgoogle/protobuf/timestamp.protoheyrenee/v1/options/auth.proto heyrenee/v1/enums/language.proto"?
Patient#

patient_id (	B??R	patientId

first_name (	R	firstName
middle_name (	R
middleName
	last_name (	RlastName#
primary_phone (	RprimaryPhone>
date_of_birth (2.google.protobuf.TimestampRdateOfBirth+
sex (2.heyrenee.v1.messages.SexRsex
email (	Remail#
address_lines	 (	RaddressLines
city
 (	Rcity
state (	Rstate
zip (	Rzip%
preferred_name (	RpreferredNameJ
marital_status (2#.heyrenee.v1.messages.MaritalStatusRmaritalStatusJ
preferred_language (2.heyrenee.v1.enums.LanguageRpreferredLanguage=
	ethnicity (2.heyrenee.v1.messages.EthnicityR	ethnicity!
mobile_phone (	RmobilePhone
other_phone (	R
otherPhone
notes (	Rnotes-
name_pronunciation (	RnamePronunciation*8
Sex
SEX_UNSPECIFIED 
SEX_MALE

SEX_FEMALE*f
MaritalStatus
MARITAL_STATUS_UNSPECIFIED 
MARITAL_STATUS_SINGLE
MARITAL_STATUS_MARRIED*?
	Ethnicity
ETHNICITY_UNSPECIFIED 
ETHNICITY_ASIAN_AMERICAN'
#ETHNICITY_BLACK_OR_AFRICAN_AMERICAN(
$ETHNICITY_WHITE_OR_EUROPEAN_AMERICAN.
*ETHNICITY_AMERICAN_INDIAN_OR_ALASKA_NATIVE1
-ETHNICITY_NATIVE_HAWAIIAN_OR_PACIFIC_ISLANDERB:Z8github.com/HeyReneeInc/proto/golang/heyrenee/v1/messagesJ?
  X

  

 

 O
	
 O
	
  )
	
 (
	
 *
L
  4@ A Patient represents a Patient subuser of the HeyRenee system.



 
/
  b" The ID of the Patient. Required.


  

  	

  

  a

  І`
8
 + The Patient's legal first name. Required.


 

 	

 
/
 " The Patient's legal middle name.


 

 	

 
7
 * The Patient's legal last name. Required.


 

 	

 
M
 @ The primary phone number for contacting the Patient. Required.


 

 	

 
5
 .( The Patient's date of birth. Required.


 

 )

 ,-
!
  The Patient's sex.


 

 	

 
5
 ( The Patient's email address. Required.


 

 	

 
@
 $3 The address of the Patient's permanent residence.


 


 

 

 "#
=
 	0 The city of the Patient's permanent residence.


 	

 		

 	
>
 
!1 The state of the Patient's permanent residence.


 
!

 
!	

 
!
A
 #4 The zip code of the Patient's permanent residence.


 #

 #	

 #
,
 % The Patient's preferred name.


 %

 %	

 %
,
 '$ The Patient's marital status.


 '

 '

 '!#
\
 )5O The Patient's preferred language. Required, must not be LANGUAGE_UNSPECIFIED.


 )

 )/

 )24
'
 + The Patient's ethnicity.


 +

 +

 +
B
 -5 The mobile phone number for contacting the Patient.


 -

 -	

 -
E
 /8 The alternate phone number for contacting the Patient.


 /

 /	

 /
=
 10 Notes about the Patient written by Concierges.


 1

 1	

 1
@
 3!3 The phonetic pronunciation of the Patient's name.


 3

 3	

 3 
-
 7 >! Sex specifies a biological sex.



 7
$
  9 Sex is not specified.


  9

  9

 ; The male sex.


 ;


 ;

 = The female sex.


 =

 =
=
A H1 MaritalStatus specifies a legal marital status.



A
/
 C!" Marital status is not specified.


 C

 C 
(
E Marital status is single.


E

E
)
G Marital status is married.


G

G
:
K X. Ethnicity specifies a legal ethnic grouping.



K
*
 M Ethnicity is not specified.


 M

 M

O Asian American.


O

O
)
Q* Black or African American.


Q%

Q()
*
S+ White or European American.


S&

S)*
0
U1# American Indian or Alaska Native.


U,

U/0
3
W4& Native Hawaiian or Pacific Islander.


W/

W23bproto3
?
heyrenee/v1/user_service.protoheyrenee.v1heyrenee/v1/options/auth.proto"heyrenee/v1/messages/patient.proto$heyrenee/v1/messages/caregiver.proto"k
CreatePatientRequest7
patient (2.heyrenee.v1.messages.PatientRpatient
password (	Rpassword"8
GetPatientRequest#

patient_id (	B??R	patientId"O
UpdatePatientRequest7
patient (2.heyrenee.v1.messages.PatientRpatient"s
CreateCaregiverRequest=
	caregiver (2.heyrenee.v1.messages.CaregiverR	caregiver
password (	Rpassword">
GetCaregiverRequest'
caregiver_id (	B??RcaregiverId"W
UpdateCaregiverRequest=
	caregiver (2.heyrenee.v1.messages.CaregiverR	caregiver2?
UserServiceS
CreatePatient!.heyrenee.v1.CreatePatientRequest.heyrenee.v1.messages.Patient" M

GetPatient.heyrenee.v1.GetPatientRequest.heyrenee.v1.messages.Patient" S
UpdatePatient!.heyrenee.v1.UpdatePatientRequest.heyrenee.v1.messages.Patient" Y
CreateCaregiver#.heyrenee.v1.CreateCaregiverRequest.heyrenee.v1.messages.Caregiver" S
GetCaregiver .heyrenee.v1.GetCaregiverRequest.heyrenee.v1.messages.Caregiver" Y
UpdateCaregiver#.heyrenee.v1.UpdateCaregiverRequest.heyrenee.v1.messages.Caregiver" B1Z/github.com/HeyReneeInc/proto/golang/heyrenee/v1J?
  I

  

 

 F
	
 F
	
  (
	
 ,
	
 .
?
   ? UserService

 The UserService provides methods for the creation and management of HeyRenee users.

 URLs
  Scratch: heyrenee-v1-8jiz6fu6.uc.gateway.dev



 
?
  S? CreatePatient creates a new User with a Patient subuser. If the request is unauthenticated, a standalone user with
 a Patient subuser will be created. If the request is authenticated by a Caregiver or Concierge subuser then the new
 Patient subuser will be created with a PatientCaregiver or a PatientConcierge for the authenticated subuser.


  

  (

  3O
@
 M3 GetPatient retrieves a specified Patient subuser.


 

 "

 -I
A
 S4 UpdatePatient updates a specified Patient subuser.


 

 (

 3O
K
 Y> CreateCaregiver creates a new User with a Caregiver subuser.


 

 ,

 7U
D
 S7 GetCaregiver retrieves a specified Caregiver subuser.


 

 &

 1O
E
 Y8 UpdateCaregiver updates a specified Caregiver subuser.


 

 ,

 7U
0
 # ($ Request message for CreatePatient.



 #
/
  %+" The Patient to create. Required.


  %

  %&

  %)*
\
 'O The Patient's password. If not provided, a random password will be generated.


 '

 '	

 '
-
+ .! Request message for GetPatient.



+
;
 -b. The ID of the Patient to retrieve. Required.


 -

 -	

 -

 -a

 І-`
0
1 4$ Request message for UpdatePatient.



1
/
 3+" The Patient to update. Required.


 3

 3&

 3)*
2
8 =& Request message for CreateCaregiver.



8
1
 :/$ The Caregiver to create. Required.


 : 

 :!*

 :-.
f
<Y Hash of the Caregiver's password. If not provided, a random password will be generated.


<

<	

<
/
@ C# Request message for GetCaregiver.



@
=
 Bf0 The ID of the Caregiver to retrieve. Required.


 B

 B	

 B

 Be

 ІBd
2
F I& Request message for UpdateCaregiver.



F
1
 H/$ The Caregiver to update. Required.


 H 

 H!*

 H-.bproto3
?
$heyrenee/v1/messages/concierge.protoheyrenee.v1.messagesheyrenee/v1/options/auth.proto"?
	Concierge'
concierge_id (	B??RconciergeId

first_name (	R	firstName
	last_name (	RlastName
email (	RemailP
concierge_status (2%.heyrenee.v1.messages.ConciergeStatusRconciergeStatus*o
ConciergeStatus 
CONCIERGE_STATUS_UNSPECIFIED 
CONCIERGE_STATUS_ACTIVE
CONCIERGE_STATUS_INACTIVEB:Z8github.com/HeyReneeInc/proto/golang/heyrenee/v1/messagesJ?
  

  

 

 O
	
 O
	
  (
?
 
 ~ A Concierge represents a subuser that is responsible for coordinating the care of Patients they are a
 PatientConcierge for.



 

1
  f$ The ID of the Concierge. Required.


  

  	

  

  e

  Іd
?
 2 The legal first name of the Concierge. Required.


 

 	

 
>
 1 The legal last name of the Concierge. Required.


 

 	

 
H
 ; The email address for contacting the Concierge. Required.


 

 	

 
5
 '( The status of the Concierge. Required.


 

 "

 %&
B
  6 ConciergeStatus specifies the status of a Concierge.



 
:
  #- The status of the Concierge is unspecified.


  

  !"
1
 $ The Concierge is currently active.


 

 
3
  & The Concierge is currently inactive.


 

 bproto3