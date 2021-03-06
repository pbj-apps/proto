// source: heyrenee/v1/messages/patient_health_questionnaire.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var heyrenee_v1_options_auth_pb = require('../../../heyrenee/v1/options/auth_pb.js');
goog.object.extend(proto, heyrenee_v1_options_auth_pb);
var heyrenee_v1_enums_language_pb = require('../../../heyrenee/v1/enums/language_pb.js');
goog.object.extend(proto, heyrenee_v1_enums_language_pb);
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
goog.object.extend(proto, google_protobuf_timestamp_pb);
goog.exportSymbol('proto.heyrenee.v1.messages.AnnualFluVaccineStatus', null, global);
goog.exportSymbol('proto.heyrenee.v1.messages.AnnualWellnessVisitStatus', null, global);
goog.exportSymbol('proto.heyrenee.v1.messages.ColorectalCancerScreenStatus', null, global);
goog.exportSymbol('proto.heyrenee.v1.messages.ColorectalCancerScreenType', null, global);
goog.exportSymbol('proto.heyrenee.v1.messages.CovidVaccineStatus', null, global);
goog.exportSymbol('proto.heyrenee.v1.messages.DailyMedicationDifficulty', null, global);
goog.exportSymbol('proto.heyrenee.v1.messages.MammogramStatus', null, global);
goog.exportSymbol('proto.heyrenee.v1.messages.MedicationCopayDifficulty', null, global);
goog.exportSymbol('proto.heyrenee.v1.messages.MedicationPaymentDifficulty', null, global);
goog.exportSymbol('proto.heyrenee.v1.messages.PatientHealthQuestionnaire', null, global);
goog.exportSymbol('proto.heyrenee.v1.messages.PillSystemUsage', null, global);
goog.exportSymbol('proto.heyrenee.v1.messages.PneumococcalVaccineStatus', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.heyrenee.v1.messages.PatientHealthQuestionnaire, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.heyrenee.v1.messages.PatientHealthQuestionnaire.displayName = 'proto.heyrenee.v1.messages.PatientHealthQuestionnaire';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.toObject = function(opt_includeInstance) {
  return proto.heyrenee.v1.messages.PatientHealthQuestionnaire.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.heyrenee.v1.messages.PatientHealthQuestionnaire} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.toObject = function(includeInstance, msg) {
  var f, obj = {
    patientId: jspb.Message.getFieldWithDefault(msg, 1, ""),
    patientHealthQuestionnaireId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    questionnaireLanguage: jspb.Message.getFieldWithDefault(msg, 3, 0),
    dateTimeAnswered: (f = msg.getDateTimeAnswered()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    annualWellnessVisitStatus: jspb.Message.getFieldWithDefault(msg, 5, 0),
    lastAnnualWellnessVisitDate: (f = msg.getLastAnnualWellnessVisitDate()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    mammogramStatus: jspb.Message.getFieldWithDefault(msg, 7, 0),
    lastMammogramDate: (f = msg.getLastMammogramDate()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    colorectalCancerScreenStatus: jspb.Message.getFieldWithDefault(msg, 9, 0),
    lastColorectalCancerScreenDate: (f = msg.getLastColorectalCancerScreenDate()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    lastColorectalCancerScreenType: jspb.Message.getFieldWithDefault(msg, 11, 0),
    annualFluVaccineStatus: jspb.Message.getFieldWithDefault(msg, 12, 0),
    lastAnnualFluVaccineDate: (f = msg.getLastAnnualFluVaccineDate()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    pneumococcalVaccineStatus: jspb.Message.getFieldWithDefault(msg, 14, 0),
    covidVaccineStatus: jspb.Message.getFieldWithDefault(msg, 15, 0),
    lastCovidVaccineDate: (f = msg.getLastCovidVaccineDate()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    pillSystemUsage: jspb.Message.getFieldWithDefault(msg, 17, 0),
    dailyMedicationDifficulty: jspb.Message.getFieldWithDefault(msg, 18, 0),
    medicationCopayDifficulty: jspb.Message.getFieldWithDefault(msg, 19, 0),
    medicationPaymentDifficulty: jspb.Message.getFieldWithDefault(msg, 20, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.heyrenee.v1.messages.PatientHealthQuestionnaire}
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.heyrenee.v1.messages.PatientHealthQuestionnaire;
  return proto.heyrenee.v1.messages.PatientHealthQuestionnaire.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.heyrenee.v1.messages.PatientHealthQuestionnaire} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.heyrenee.v1.messages.PatientHealthQuestionnaire}
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setPatientId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setPatientHealthQuestionnaireId(value);
      break;
    case 3:
      var value = /** @type {!proto.heyrenee.v1.enums.Language} */ (reader.readEnum());
      msg.setQuestionnaireLanguage(value);
      break;
    case 4:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setDateTimeAnswered(value);
      break;
    case 5:
      var value = /** @type {!proto.heyrenee.v1.messages.AnnualWellnessVisitStatus} */ (reader.readEnum());
      msg.setAnnualWellnessVisitStatus(value);
      break;
    case 6:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setLastAnnualWellnessVisitDate(value);
      break;
    case 7:
      var value = /** @type {!proto.heyrenee.v1.messages.MammogramStatus} */ (reader.readEnum());
      msg.setMammogramStatus(value);
      break;
    case 8:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setLastMammogramDate(value);
      break;
    case 9:
      var value = /** @type {!proto.heyrenee.v1.messages.ColorectalCancerScreenStatus} */ (reader.readEnum());
      msg.setColorectalCancerScreenStatus(value);
      break;
    case 10:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setLastColorectalCancerScreenDate(value);
      break;
    case 11:
      var value = /** @type {!proto.heyrenee.v1.messages.ColorectalCancerScreenType} */ (reader.readEnum());
      msg.setLastColorectalCancerScreenType(value);
      break;
    case 12:
      var value = /** @type {!proto.heyrenee.v1.messages.AnnualFluVaccineStatus} */ (reader.readEnum());
      msg.setAnnualFluVaccineStatus(value);
      break;
    case 13:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setLastAnnualFluVaccineDate(value);
      break;
    case 14:
      var value = /** @type {!proto.heyrenee.v1.messages.PneumococcalVaccineStatus} */ (reader.readEnum());
      msg.setPneumococcalVaccineStatus(value);
      break;
    case 15:
      var value = /** @type {!proto.heyrenee.v1.messages.CovidVaccineStatus} */ (reader.readEnum());
      msg.setCovidVaccineStatus(value);
      break;
    case 16:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setLastCovidVaccineDate(value);
      break;
    case 17:
      var value = /** @type {!proto.heyrenee.v1.messages.PillSystemUsage} */ (reader.readEnum());
      msg.setPillSystemUsage(value);
      break;
    case 18:
      var value = /** @type {!proto.heyrenee.v1.messages.DailyMedicationDifficulty} */ (reader.readEnum());
      msg.setDailyMedicationDifficulty(value);
      break;
    case 19:
      var value = /** @type {!proto.heyrenee.v1.messages.MedicationCopayDifficulty} */ (reader.readEnum());
      msg.setMedicationCopayDifficulty(value);
      break;
    case 20:
      var value = /** @type {!proto.heyrenee.v1.messages.MedicationPaymentDifficulty} */ (reader.readEnum());
      msg.setMedicationPaymentDifficulty(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.heyrenee.v1.messages.PatientHealthQuestionnaire.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.heyrenee.v1.messages.PatientHealthQuestionnaire} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getPatientId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getPatientHealthQuestionnaireId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getQuestionnaireLanguage();
  if (f !== 0.0) {
    writer.writeEnum(
      3,
      f
    );
  }
  f = message.getDateTimeAnswered();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getAnnualWellnessVisitStatus();
  if (f !== 0.0) {
    writer.writeEnum(
      5,
      f
    );
  }
  f = message.getLastAnnualWellnessVisitDate();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getMammogramStatus();
  if (f !== 0.0) {
    writer.writeEnum(
      7,
      f
    );
  }
  f = message.getLastMammogramDate();
  if (f != null) {
    writer.writeMessage(
      8,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getColorectalCancerScreenStatus();
  if (f !== 0.0) {
    writer.writeEnum(
      9,
      f
    );
  }
  f = message.getLastColorectalCancerScreenDate();
  if (f != null) {
    writer.writeMessage(
      10,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getLastColorectalCancerScreenType();
  if (f !== 0.0) {
    writer.writeEnum(
      11,
      f
    );
  }
  f = message.getAnnualFluVaccineStatus();
  if (f !== 0.0) {
    writer.writeEnum(
      12,
      f
    );
  }
  f = message.getLastAnnualFluVaccineDate();
  if (f != null) {
    writer.writeMessage(
      13,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getPneumococcalVaccineStatus();
  if (f !== 0.0) {
    writer.writeEnum(
      14,
      f
    );
  }
  f = message.getCovidVaccineStatus();
  if (f !== 0.0) {
    writer.writeEnum(
      15,
      f
    );
  }
  f = message.getLastCovidVaccineDate();
  if (f != null) {
    writer.writeMessage(
      16,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getPillSystemUsage();
  if (f !== 0.0) {
    writer.writeEnum(
      17,
      f
    );
  }
  f = message.getDailyMedicationDifficulty();
  if (f !== 0.0) {
    writer.writeEnum(
      18,
      f
    );
  }
  f = message.getMedicationCopayDifficulty();
  if (f !== 0.0) {
    writer.writeEnum(
      19,
      f
    );
  }
  f = message.getMedicationPaymentDifficulty();
  if (f !== 0.0) {
    writer.writeEnum(
      20,
      f
    );
  }
};


/**
 * optional string patient_id = 1;
 * @return {string}
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.getPatientId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.messages.PatientHealthQuestionnaire} returns this
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.setPatientId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string patient_health_questionnaire_id = 2;
 * @return {string}
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.getPatientHealthQuestionnaireId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.messages.PatientHealthQuestionnaire} returns this
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.setPatientHealthQuestionnaireId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional heyrenee.v1.enums.Language questionnaire_language = 3;
 * @return {!proto.heyrenee.v1.enums.Language}
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.getQuestionnaireLanguage = function() {
  return /** @type {!proto.heyrenee.v1.enums.Language} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {!proto.heyrenee.v1.enums.Language} value
 * @return {!proto.heyrenee.v1.messages.PatientHealthQuestionnaire} returns this
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.setQuestionnaireLanguage = function(value) {
  return jspb.Message.setProto3EnumField(this, 3, value);
};


/**
 * optional google.protobuf.Timestamp date_time_answered = 4;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.getDateTimeAnswered = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 4));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.heyrenee.v1.messages.PatientHealthQuestionnaire} returns this
*/
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.setDateTimeAnswered = function(value) {
  return jspb.Message.setWrapperField(this, 4, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.heyrenee.v1.messages.PatientHealthQuestionnaire} returns this
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.clearDateTimeAnswered = function() {
  return this.setDateTimeAnswered(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.hasDateTimeAnswered = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional AnnualWellnessVisitStatus annual_wellness_visit_status = 5;
 * @return {!proto.heyrenee.v1.messages.AnnualWellnessVisitStatus}
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.getAnnualWellnessVisitStatus = function() {
  return /** @type {!proto.heyrenee.v1.messages.AnnualWellnessVisitStatus} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/**
 * @param {!proto.heyrenee.v1.messages.AnnualWellnessVisitStatus} value
 * @return {!proto.heyrenee.v1.messages.PatientHealthQuestionnaire} returns this
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.setAnnualWellnessVisitStatus = function(value) {
  return jspb.Message.setProto3EnumField(this, 5, value);
};


/**
 * optional google.protobuf.Timestamp last_annual_wellness_visit_date = 6;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.getLastAnnualWellnessVisitDate = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 6));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.heyrenee.v1.messages.PatientHealthQuestionnaire} returns this
*/
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.setLastAnnualWellnessVisitDate = function(value) {
  return jspb.Message.setWrapperField(this, 6, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.heyrenee.v1.messages.PatientHealthQuestionnaire} returns this
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.clearLastAnnualWellnessVisitDate = function() {
  return this.setLastAnnualWellnessVisitDate(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.hasLastAnnualWellnessVisitDate = function() {
  return jspb.Message.getField(this, 6) != null;
};


/**
 * optional MammogramStatus mammogram_status = 7;
 * @return {!proto.heyrenee.v1.messages.MammogramStatus}
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.getMammogramStatus = function() {
  return /** @type {!proto.heyrenee.v1.messages.MammogramStatus} */ (jspb.Message.getFieldWithDefault(this, 7, 0));
};


/**
 * @param {!proto.heyrenee.v1.messages.MammogramStatus} value
 * @return {!proto.heyrenee.v1.messages.PatientHealthQuestionnaire} returns this
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.setMammogramStatus = function(value) {
  return jspb.Message.setProto3EnumField(this, 7, value);
};


/**
 * optional google.protobuf.Timestamp last_mammogram_date = 8;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.getLastMammogramDate = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 8));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.heyrenee.v1.messages.PatientHealthQuestionnaire} returns this
*/
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.setLastMammogramDate = function(value) {
  return jspb.Message.setWrapperField(this, 8, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.heyrenee.v1.messages.PatientHealthQuestionnaire} returns this
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.clearLastMammogramDate = function() {
  return this.setLastMammogramDate(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.hasLastMammogramDate = function() {
  return jspb.Message.getField(this, 8) != null;
};


/**
 * optional ColorectalCancerScreenStatus colorectal_cancer_screen_status = 9;
 * @return {!proto.heyrenee.v1.messages.ColorectalCancerScreenStatus}
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.getColorectalCancerScreenStatus = function() {
  return /** @type {!proto.heyrenee.v1.messages.ColorectalCancerScreenStatus} */ (jspb.Message.getFieldWithDefault(this, 9, 0));
};


/**
 * @param {!proto.heyrenee.v1.messages.ColorectalCancerScreenStatus} value
 * @return {!proto.heyrenee.v1.messages.PatientHealthQuestionnaire} returns this
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.setColorectalCancerScreenStatus = function(value) {
  return jspb.Message.setProto3EnumField(this, 9, value);
};


/**
 * optional google.protobuf.Timestamp last_colorectal_cancer_screen_date = 10;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.getLastColorectalCancerScreenDate = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 10));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.heyrenee.v1.messages.PatientHealthQuestionnaire} returns this
*/
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.setLastColorectalCancerScreenDate = function(value) {
  return jspb.Message.setWrapperField(this, 10, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.heyrenee.v1.messages.PatientHealthQuestionnaire} returns this
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.clearLastColorectalCancerScreenDate = function() {
  return this.setLastColorectalCancerScreenDate(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.hasLastColorectalCancerScreenDate = function() {
  return jspb.Message.getField(this, 10) != null;
};


/**
 * optional ColorectalCancerScreenType last_colorectal_cancer_screen_type = 11;
 * @return {!proto.heyrenee.v1.messages.ColorectalCancerScreenType}
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.getLastColorectalCancerScreenType = function() {
  return /** @type {!proto.heyrenee.v1.messages.ColorectalCancerScreenType} */ (jspb.Message.getFieldWithDefault(this, 11, 0));
};


/**
 * @param {!proto.heyrenee.v1.messages.ColorectalCancerScreenType} value
 * @return {!proto.heyrenee.v1.messages.PatientHealthQuestionnaire} returns this
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.setLastColorectalCancerScreenType = function(value) {
  return jspb.Message.setProto3EnumField(this, 11, value);
};


/**
 * optional AnnualFluVaccineStatus annual_flu_vaccine_status = 12;
 * @return {!proto.heyrenee.v1.messages.AnnualFluVaccineStatus}
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.getAnnualFluVaccineStatus = function() {
  return /** @type {!proto.heyrenee.v1.messages.AnnualFluVaccineStatus} */ (jspb.Message.getFieldWithDefault(this, 12, 0));
};


/**
 * @param {!proto.heyrenee.v1.messages.AnnualFluVaccineStatus} value
 * @return {!proto.heyrenee.v1.messages.PatientHealthQuestionnaire} returns this
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.setAnnualFluVaccineStatus = function(value) {
  return jspb.Message.setProto3EnumField(this, 12, value);
};


/**
 * optional google.protobuf.Timestamp last_annual_flu_vaccine_date = 13;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.getLastAnnualFluVaccineDate = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 13));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.heyrenee.v1.messages.PatientHealthQuestionnaire} returns this
*/
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.setLastAnnualFluVaccineDate = function(value) {
  return jspb.Message.setWrapperField(this, 13, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.heyrenee.v1.messages.PatientHealthQuestionnaire} returns this
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.clearLastAnnualFluVaccineDate = function() {
  return this.setLastAnnualFluVaccineDate(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.hasLastAnnualFluVaccineDate = function() {
  return jspb.Message.getField(this, 13) != null;
};


/**
 * optional PneumococcalVaccineStatus pneumococcal_vaccine_status = 14;
 * @return {!proto.heyrenee.v1.messages.PneumococcalVaccineStatus}
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.getPneumococcalVaccineStatus = function() {
  return /** @type {!proto.heyrenee.v1.messages.PneumococcalVaccineStatus} */ (jspb.Message.getFieldWithDefault(this, 14, 0));
};


/**
 * @param {!proto.heyrenee.v1.messages.PneumococcalVaccineStatus} value
 * @return {!proto.heyrenee.v1.messages.PatientHealthQuestionnaire} returns this
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.setPneumococcalVaccineStatus = function(value) {
  return jspb.Message.setProto3EnumField(this, 14, value);
};


/**
 * optional CovidVaccineStatus covid_vaccine_status = 15;
 * @return {!proto.heyrenee.v1.messages.CovidVaccineStatus}
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.getCovidVaccineStatus = function() {
  return /** @type {!proto.heyrenee.v1.messages.CovidVaccineStatus} */ (jspb.Message.getFieldWithDefault(this, 15, 0));
};


/**
 * @param {!proto.heyrenee.v1.messages.CovidVaccineStatus} value
 * @return {!proto.heyrenee.v1.messages.PatientHealthQuestionnaire} returns this
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.setCovidVaccineStatus = function(value) {
  return jspb.Message.setProto3EnumField(this, 15, value);
};


/**
 * optional google.protobuf.Timestamp last_covid_vaccine_date = 16;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.getLastCovidVaccineDate = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 16));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.heyrenee.v1.messages.PatientHealthQuestionnaire} returns this
*/
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.setLastCovidVaccineDate = function(value) {
  return jspb.Message.setWrapperField(this, 16, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.heyrenee.v1.messages.PatientHealthQuestionnaire} returns this
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.clearLastCovidVaccineDate = function() {
  return this.setLastCovidVaccineDate(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.hasLastCovidVaccineDate = function() {
  return jspb.Message.getField(this, 16) != null;
};


/**
 * optional PillSystemUsage pill_system_usage = 17;
 * @return {!proto.heyrenee.v1.messages.PillSystemUsage}
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.getPillSystemUsage = function() {
  return /** @type {!proto.heyrenee.v1.messages.PillSystemUsage} */ (jspb.Message.getFieldWithDefault(this, 17, 0));
};


/**
 * @param {!proto.heyrenee.v1.messages.PillSystemUsage} value
 * @return {!proto.heyrenee.v1.messages.PatientHealthQuestionnaire} returns this
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.setPillSystemUsage = function(value) {
  return jspb.Message.setProto3EnumField(this, 17, value);
};


/**
 * optional DailyMedicationDifficulty daily_medication_difficulty = 18;
 * @return {!proto.heyrenee.v1.messages.DailyMedicationDifficulty}
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.getDailyMedicationDifficulty = function() {
  return /** @type {!proto.heyrenee.v1.messages.DailyMedicationDifficulty} */ (jspb.Message.getFieldWithDefault(this, 18, 0));
};


/**
 * @param {!proto.heyrenee.v1.messages.DailyMedicationDifficulty} value
 * @return {!proto.heyrenee.v1.messages.PatientHealthQuestionnaire} returns this
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.setDailyMedicationDifficulty = function(value) {
  return jspb.Message.setProto3EnumField(this, 18, value);
};


/**
 * optional MedicationCopayDifficulty medication_copay_difficulty = 19;
 * @return {!proto.heyrenee.v1.messages.MedicationCopayDifficulty}
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.getMedicationCopayDifficulty = function() {
  return /** @type {!proto.heyrenee.v1.messages.MedicationCopayDifficulty} */ (jspb.Message.getFieldWithDefault(this, 19, 0));
};


/**
 * @param {!proto.heyrenee.v1.messages.MedicationCopayDifficulty} value
 * @return {!proto.heyrenee.v1.messages.PatientHealthQuestionnaire} returns this
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.setMedicationCopayDifficulty = function(value) {
  return jspb.Message.setProto3EnumField(this, 19, value);
};


/**
 * optional MedicationPaymentDifficulty medication_payment_difficulty = 20;
 * @return {!proto.heyrenee.v1.messages.MedicationPaymentDifficulty}
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.getMedicationPaymentDifficulty = function() {
  return /** @type {!proto.heyrenee.v1.messages.MedicationPaymentDifficulty} */ (jspb.Message.getFieldWithDefault(this, 20, 0));
};


/**
 * @param {!proto.heyrenee.v1.messages.MedicationPaymentDifficulty} value
 * @return {!proto.heyrenee.v1.messages.PatientHealthQuestionnaire} returns this
 */
proto.heyrenee.v1.messages.PatientHealthQuestionnaire.prototype.setMedicationPaymentDifficulty = function(value) {
  return jspb.Message.setProto3EnumField(this, 20, value);
};


/**
 * @enum {number}
 */
proto.heyrenee.v1.messages.AnnualWellnessVisitStatus = {
  ANNUAL_WELLNESS_VISIT_STATUS_UNSPECIFIED: 0,
  ANNUAL_WELLNESS_VISIT_STATUS_COMPLETED: 1,
  ANNUAL_WELLNESS_VISIT_STATUS_NOT_COMPLETED: 2
};

/**
 * @enum {number}
 */
proto.heyrenee.v1.messages.MammogramStatus = {
  MAMMOGRAM_STATUS_UNSPECIFIED: 0,
  MAMMOGRAM_STATUS_COMPLETED: 1,
  MAMMOGRAM_STATUS_NOT_COMPLETED: 2
};

/**
 * @enum {number}
 */
proto.heyrenee.v1.messages.ColorectalCancerScreenStatus = {
  COLORECTAL_CANCER_SCREEN_STATUS_UNSPECIFIED: 0,
  COLORECTAL_CANCER_SCREEN_STATUS_COMPLETED: 1,
  COLORECTAL_CANCER_SCREEN_STATUS_NOT_COMPLETED: 2
};

/**
 * @enum {number}
 */
proto.heyrenee.v1.messages.ColorectalCancerScreenType = {
  COLORECTAL_CANCER_SCREEN_TYPE_UNSPECIFIED: 0,
  COLORECTAL_CANCER_SCREEN_TYPE_FIT: 1,
  COLORECTAL_CANCER_SCREEN_TYPE_COLOGUARD: 2,
  COLORECTAL_CANCER_SCREEN_TYPE_COLONOSCOPY: 3
};

/**
 * @enum {number}
 */
proto.heyrenee.v1.messages.AnnualFluVaccineStatus = {
  ANNUAL_FLU_VACCINE_STATUS_UNSPECIFIED: 0,
  ANNUAL_FLU_VACCINE_STATUS_COMPLETED: 1,
  ANNUAL_FLU_VACCINE_STATUS_NOT_COMPLETED: 2
};

/**
 * @enum {number}
 */
proto.heyrenee.v1.messages.PneumococcalVaccineStatus = {
  PNEUMOCOCCAL_VACCINE_STATUS_UNSPECIFIED: 0,
  PNEUMOCOCCAL_VACCINE_STATUS_COMPLETED: 1,
  PNEUMOCOCCAL_VACCINE_STATUS_ONLY_ONE: 2,
  PNEUMOCOCCAL_VACCINE_STATUS_NOT_COMPLETED: 3
};

/**
 * @enum {number}
 */
proto.heyrenee.v1.messages.CovidVaccineStatus = {
  COVID_VACCINE_STATUS_UNSPECIFIED: 0,
  COVID_VACCINE_STATUS_COMPLETED: 1,
  COVID_VACCINE_STATUS_NOT_COMPLETED: 2
};

/**
 * @enum {number}
 */
proto.heyrenee.v1.messages.PillSystemUsage = {
  PILL_SYSTEM_USAGE_NOT_SPECIFIED: 0,
  PILL_SYSTEM_USAGE_USES: 1,
  PILL_SYSTEM_USAGE_DOES_NOT_USE: 2
};

/**
 * @enum {number}
 */
proto.heyrenee.v1.messages.DailyMedicationDifficulty = {
  DAILY_MEDICATION_DIFFICULTY_NOT_SPECIFIED: 0,
  DAILY_MEDICATION_DIFFICULTY_HAS_DIFFICULTY: 1,
  DAILY_MEDICATION_DIFFICULTY_NO_DIFFICULTY: 2
};

/**
 * @enum {number}
 */
proto.heyrenee.v1.messages.MedicationCopayDifficulty = {
  MEDICATION_COPAY_DIFFICULTY_NOT_SPECIFIED: 0,
  MEDICATION_COPAY_DIFFICULTY_HAS_DIFFICULTY: 1,
  MEDICATION_COPAY_DIFFICULTY_NO_DIFFICULTY: 2
};

/**
 * @enum {number}
 */
proto.heyrenee.v1.messages.MedicationPaymentDifficulty = {
  MEDICATION_PAYMENT_DIFFICULTY_NOT_SPECIFIED: 0,
  MEDICATION_PAYMENT_DIFFICULTY_SKIPS_BECAUSE_CANT_AFFORD: 1,
  MEDICATION_PAYMENT_DIFFICULTY_NO_DIFFICULTY: 2
};

goog.object.extend(exports, proto.heyrenee.v1.messages);
