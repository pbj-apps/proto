// package: heyrenee.v1.messages
// file: heyrenee/v1/messages/patient_health_questionnaire.proto

import * as jspb from "google-protobuf";
import * as heyrenee_v1_options_auth_pb from "../../../heyrenee/v1/options/auth_pb";
import * as heyrenee_v1_enums_language_pb from "../../../heyrenee/v1/enums/language_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class PatientHealthQuestionnaire extends jspb.Message {
  getPatientId(): string;
  setPatientId(value: string): void;

  getPatientHealthQuestionnaireId(): string;
  setPatientHealthQuestionnaireId(value: string): void;

  getQuestionnaireLanguage(): heyrenee_v1_enums_language_pb.LanguageMap[keyof heyrenee_v1_enums_language_pb.LanguageMap];
  setQuestionnaireLanguage(value: heyrenee_v1_enums_language_pb.LanguageMap[keyof heyrenee_v1_enums_language_pb.LanguageMap]): void;

  hasDateTimeAnswered(): boolean;
  clearDateTimeAnswered(): void;
  getDateTimeAnswered(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDateTimeAnswered(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getAnnualWellnessVisitStatus(): AnnualWellnessVisitStatusMap[keyof AnnualWellnessVisitStatusMap];
  setAnnualWellnessVisitStatus(value: AnnualWellnessVisitStatusMap[keyof AnnualWellnessVisitStatusMap]): void;

  hasLastAnnualWellnessVisitDate(): boolean;
  clearLastAnnualWellnessVisitDate(): void;
  getLastAnnualWellnessVisitDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setLastAnnualWellnessVisitDate(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getMammogramStatus(): MammogramStatusMap[keyof MammogramStatusMap];
  setMammogramStatus(value: MammogramStatusMap[keyof MammogramStatusMap]): void;

  hasLastMammogramDate(): boolean;
  clearLastMammogramDate(): void;
  getLastMammogramDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setLastMammogramDate(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getColorectalCancerScreenStatus(): ColorectalCancerScreenStatusMap[keyof ColorectalCancerScreenStatusMap];
  setColorectalCancerScreenStatus(value: ColorectalCancerScreenStatusMap[keyof ColorectalCancerScreenStatusMap]): void;

  hasLastColorectalCancerScreenDate(): boolean;
  clearLastColorectalCancerScreenDate(): void;
  getLastColorectalCancerScreenDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setLastColorectalCancerScreenDate(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getLastColorectalCancerScreenType(): ColorectalCancerScreenTypeMap[keyof ColorectalCancerScreenTypeMap];
  setLastColorectalCancerScreenType(value: ColorectalCancerScreenTypeMap[keyof ColorectalCancerScreenTypeMap]): void;

  getAnnualFluVaccineStatus(): AnnualFluVaccineStatusMap[keyof AnnualFluVaccineStatusMap];
  setAnnualFluVaccineStatus(value: AnnualFluVaccineStatusMap[keyof AnnualFluVaccineStatusMap]): void;

  hasLastAnnualFluVaccineDate(): boolean;
  clearLastAnnualFluVaccineDate(): void;
  getLastAnnualFluVaccineDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setLastAnnualFluVaccineDate(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getPneumococcalVaccineStatus(): PneumococcalVaccineStatusMap[keyof PneumococcalVaccineStatusMap];
  setPneumococcalVaccineStatus(value: PneumococcalVaccineStatusMap[keyof PneumococcalVaccineStatusMap]): void;

  getCovidVaccineStatus(): CovidVaccineStatusMap[keyof CovidVaccineStatusMap];
  setCovidVaccineStatus(value: CovidVaccineStatusMap[keyof CovidVaccineStatusMap]): void;

  hasLastCovidVaccineDate(): boolean;
  clearLastCovidVaccineDate(): void;
  getLastCovidVaccineDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setLastCovidVaccineDate(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getPillSystemUsage(): PillSystemUsageMap[keyof PillSystemUsageMap];
  setPillSystemUsage(value: PillSystemUsageMap[keyof PillSystemUsageMap]): void;

  getDailyMedicationDifficulty(): DailyMedicationDifficultyMap[keyof DailyMedicationDifficultyMap];
  setDailyMedicationDifficulty(value: DailyMedicationDifficultyMap[keyof DailyMedicationDifficultyMap]): void;

  getMedicationCopayDifficulty(): MedicationCopayDifficultyMap[keyof MedicationCopayDifficultyMap];
  setMedicationCopayDifficulty(value: MedicationCopayDifficultyMap[keyof MedicationCopayDifficultyMap]): void;

  getMedicationPaymentDifficulty(): MedicationPaymentDifficultyMap[keyof MedicationPaymentDifficultyMap];
  setMedicationPaymentDifficulty(value: MedicationPaymentDifficultyMap[keyof MedicationPaymentDifficultyMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PatientHealthQuestionnaire.AsObject;
  static toObject(includeInstance: boolean, msg: PatientHealthQuestionnaire): PatientHealthQuestionnaire.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PatientHealthQuestionnaire, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PatientHealthQuestionnaire;
  static deserializeBinaryFromReader(message: PatientHealthQuestionnaire, reader: jspb.BinaryReader): PatientHealthQuestionnaire;
}

export namespace PatientHealthQuestionnaire {
  export type AsObject = {
    patientId: string,
    patientHealthQuestionnaireId: string,
    questionnaireLanguage: heyrenee_v1_enums_language_pb.LanguageMap[keyof heyrenee_v1_enums_language_pb.LanguageMap],
    dateTimeAnswered?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    annualWellnessVisitStatus: AnnualWellnessVisitStatusMap[keyof AnnualWellnessVisitStatusMap],
    lastAnnualWellnessVisitDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    mammogramStatus: MammogramStatusMap[keyof MammogramStatusMap],
    lastMammogramDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    colorectalCancerScreenStatus: ColorectalCancerScreenStatusMap[keyof ColorectalCancerScreenStatusMap],
    lastColorectalCancerScreenDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    lastColorectalCancerScreenType: ColorectalCancerScreenTypeMap[keyof ColorectalCancerScreenTypeMap],
    annualFluVaccineStatus: AnnualFluVaccineStatusMap[keyof AnnualFluVaccineStatusMap],
    lastAnnualFluVaccineDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    pneumococcalVaccineStatus: PneumococcalVaccineStatusMap[keyof PneumococcalVaccineStatusMap],
    covidVaccineStatus: CovidVaccineStatusMap[keyof CovidVaccineStatusMap],
    lastCovidVaccineDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    pillSystemUsage: PillSystemUsageMap[keyof PillSystemUsageMap],
    dailyMedicationDifficulty: DailyMedicationDifficultyMap[keyof DailyMedicationDifficultyMap],
    medicationCopayDifficulty: MedicationCopayDifficultyMap[keyof MedicationCopayDifficultyMap],
    medicationPaymentDifficulty: MedicationPaymentDifficultyMap[keyof MedicationPaymentDifficultyMap],
  }
}

export interface AnnualWellnessVisitStatusMap {
  ANNUAL_WELLNESS_VISIT_STATUS_UNSPECIFIED: 0;
  ANNUAL_WELLNESS_VISIT_STATUS_COMPLETED: 1;
  ANNUAL_WELLNESS_VISIT_STATUS_NOT_COMPLETED: 2;
}

export const AnnualWellnessVisitStatus: AnnualWellnessVisitStatusMap;

export interface MammogramStatusMap {
  MAMMOGRAM_STATUS_UNSPECIFIED: 0;
  MAMMOGRAM_STATUS_COMPLETED: 1;
  MAMMOGRAM_STATUS_NOT_COMPLETED: 2;
}

export const MammogramStatus: MammogramStatusMap;

export interface ColorectalCancerScreenStatusMap {
  COLORECTAL_CANCER_SCREEN_STATUS_UNSPECIFIED: 0;
  COLORECTAL_CANCER_SCREEN_STATUS_COMPLETED: 1;
  COLORECTAL_CANCER_SCREEN_STATUS_NOT_COMPLETED: 2;
}

export const ColorectalCancerScreenStatus: ColorectalCancerScreenStatusMap;

export interface ColorectalCancerScreenTypeMap {
  COLORECTAL_CANCER_SCREEN_TYPE_UNSPECIFIED: 0;
  COLORECTAL_CANCER_SCREEN_TYPE_FIT: 1;
  COLORECTAL_CANCER_SCREEN_TYPE_COLOGUARD: 2;
  COLORECTAL_CANCER_SCREEN_TYPE_COLONOSCOPY: 3;
}

export const ColorectalCancerScreenType: ColorectalCancerScreenTypeMap;

export interface AnnualFluVaccineStatusMap {
  ANNUAL_FLU_VACCINE_STATUS_UNSPECIFIED: 0;
  ANNUAL_FLU_VACCINE_STATUS_COMPLETED: 1;
  ANNUAL_FLU_VACCINE_STATUS_NOT_COMPLETED: 2;
}

export const AnnualFluVaccineStatus: AnnualFluVaccineStatusMap;

export interface PneumococcalVaccineStatusMap {
  PNEUMOCOCCAL_VACCINE_STATUS_UNSPECIFIED: 0;
  PNEUMOCOCCAL_VACCINE_STATUS_COMPLETED: 1;
  PNEUMOCOCCAL_VACCINE_STATUS_ONLY_ONE: 2;
  PNEUMOCOCCAL_VACCINE_STATUS_NOT_COMPLETED: 3;
}

export const PneumococcalVaccineStatus: PneumococcalVaccineStatusMap;

export interface CovidVaccineStatusMap {
  COVID_VACCINE_STATUS_UNSPECIFIED: 0;
  COVID_VACCINE_STATUS_COMPLETED: 1;
  COVID_VACCINE_STATUS_NOT_COMPLETED: 2;
}

export const CovidVaccineStatus: CovidVaccineStatusMap;

export interface PillSystemUsageMap {
  PILL_SYSTEM_USAGE_NOT_SPECIFIED: 0;
  PILL_SYSTEM_USAGE_USES: 1;
  PILL_SYSTEM_USAGE_DOES_NOT_USE: 2;
}

export const PillSystemUsage: PillSystemUsageMap;

export interface DailyMedicationDifficultyMap {
  DAILY_MEDICATION_DIFFICULTY_NOT_SPECIFIED: 0;
  DAILY_MEDICATION_DIFFICULTY_HAS_DIFFICULTY: 1;
  DAILY_MEDICATION_DIFFICULTY_NO_DIFFICULTY: 2;
}

export const DailyMedicationDifficulty: DailyMedicationDifficultyMap;

export interface MedicationCopayDifficultyMap {
  MEDICATION_COPAY_DIFFICULTY_NOT_SPECIFIED: 0;
  MEDICATION_COPAY_DIFFICULTY_HAS_DIFFICULTY: 1;
  MEDICATION_COPAY_DIFFICULTY_NO_DIFFICULTY: 2;
}

export const MedicationCopayDifficulty: MedicationCopayDifficultyMap;

export interface MedicationPaymentDifficultyMap {
  MEDICATION_PAYMENT_DIFFICULTY_NOT_SPECIFIED: 0;
  MEDICATION_PAYMENT_DIFFICULTY_SKIPS_BECAUSE_CANT_AFFORD: 1;
  MEDICATION_PAYMENT_DIFFICULTY_NO_DIFFICULTY: 2;
}

export const MedicationPaymentDifficulty: MedicationPaymentDifficultyMap;

