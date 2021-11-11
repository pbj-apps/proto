// package: heyrenee.v1.messages
// file: heyrenee/v1/messages/patient_sdoh_questionnaire.proto

import * as jspb from "google-protobuf";
import * as heyrenee_v1_options_auth_pb from "../../../heyrenee/v1/options/auth_pb";
import * as heyrenee_v1_enums_language_pb from "../../../heyrenee/v1/enums/language_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class PatientSdohQuestionnaire extends jspb.Message {
  getPatientId(): string;
  setPatientId(value: string): void;

  getPatientSdohQuestionnaireId(): string;
  setPatientSdohQuestionnaireId(value: string): void;

  getQuestionnaireLanguage(): heyrenee_v1_enums_language_pb.LanguageMap[keyof heyrenee_v1_enums_language_pb.LanguageMap];
  setQuestionnaireLanguage(value: heyrenee_v1_enums_language_pb.LanguageMap[keyof heyrenee_v1_enums_language_pb.LanguageMap]): void;

  hasDateTimeAnswered(): boolean;
  clearDateTimeAnswered(): void;
  getDateTimeAnswered(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDateTimeAnswered(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getAmbulatoryStatus(): AmbulatoryStatusMap[keyof AmbulatoryStatusMap];
  setAmbulatoryStatus(value: AmbulatoryStatusMap[keyof AmbulatoryStatusMap]): void;

  getFallHistory(): FallHistoryMap[keyof FallHistoryMap];
  setFallHistory(value: FallHistoryMap[keyof FallHistoryMap]): void;

  getWalkingSteadiness(): WalkingSteadinessMap[keyof WalkingSteadinessMap];
  setWalkingSteadiness(value: WalkingSteadinessMap[keyof WalkingSteadinessMap]): void;

  getWalkingSupportNecessity(): WalkingSupportNecessityMap[keyof WalkingSupportNecessityMap];
  setWalkingSupportNecessity(value: WalkingSupportNecessityMap[keyof WalkingSupportNecessityMap]): void;

  getWalkingAssistanceAdvisory(): WalkingAssistanceAdvisoryMap[keyof WalkingAssistanceAdvisoryMap];
  setWalkingAssistanceAdvisory(value: WalkingAssistanceAdvisoryMap[keyof WalkingAssistanceAdvisoryMap]): void;

  getWalkingAssistanceOwnership(): WalkingAssistanceOwnershipMap[keyof WalkingAssistanceOwnershipMap];
  setWalkingAssistanceOwnership(value: WalkingAssistanceOwnershipMap[keyof WalkingAssistanceOwnershipMap]): void;

  getFallingWorry(): FallingWorryMap[keyof FallingWorryMap];
  setFallingWorry(value: FallingWorryMap[keyof FallingWorryMap]): void;

  getStandingUpAssistance(): StandingUpAssistanceMap[keyof StandingUpAssistanceMap];
  setStandingUpAssistance(value: StandingUpAssistanceMap[keyof StandingUpAssistanceMap]): void;

  getSteppingUpDifficulty(): SteppingUpDifficultyMap[keyof SteppingUpDifficultyMap];
  setSteppingUpDifficulty(value: SteppingUpDifficultyMap[keyof SteppingUpDifficultyMap]): void;

  getFeelingInFeetStatus(): FeelingInFeetStatusMap[keyof FeelingInFeetStatusMap];
  setFeelingInFeetStatus(value: FeelingInFeetStatusMap[keyof FeelingInFeetStatusMap]): void;

  getFeelingLightheadedFromMedsStatus(): FeelingLightheadedFromMedsStatusMap[keyof FeelingLightheadedFromMedsStatusMap];
  setFeelingLightheadedFromMedsStatus(value: FeelingLightheadedFromMedsStatusMap[keyof FeelingLightheadedFromMedsStatusMap]): void;

  getSmokingStatus(): SmokingStatusMap[keyof SmokingStatusMap];
  setSmokingStatus(value: SmokingStatusMap[keyof SmokingStatusMap]): void;

  getAmountSmoked(): string;
  setAmountSmoked(value: string): void;

  getQuitSmokingStatus(): QuitSmokingStatusMap[keyof QuitSmokingStatusMap];
  setQuitSmokingStatus(value: QuitSmokingStatusMap[keyof QuitSmokingStatusMap]): void;

  getAlcoholDrinkingStatus(): AlcoholDrinkingStatusMap[keyof AlcoholDrinkingStatusMap];
  setAlcoholDrinkingStatus(value: AlcoholDrinkingStatusMap[keyof AlcoholDrinkingStatusMap]): void;

  getAmountAlcoholDrank(): string;
  setAmountAlcoholDrank(value: string): void;

  getQuitDrinkingStatus(): QuitDrinkingStatusMap[keyof QuitDrinkingStatusMap];
  setQuitDrinkingStatus(value: QuitDrinkingStatusMap[keyof QuitDrinkingStatusMap]): void;

  getSchedulingDoctorsAppointmentsDifficulty(): SchedulingDoctorsAppointmentsDifficultyMap[keyof SchedulingDoctorsAppointmentsDifficultyMap];
  setSchedulingDoctorsAppointmentsDifficulty(value: SchedulingDoctorsAppointmentsDifficultyMap[keyof SchedulingDoctorsAppointmentsDifficultyMap]): void;

  getGettingToDoctorsAppointmentsDifficulty(): GettingToDoctorsAppointmentsDifficultyMap[keyof GettingToDoctorsAppointmentsDifficultyMap];
  setGettingToDoctorsAppointmentsDifficulty(value: GettingToDoctorsAppointmentsDifficultyMap[keyof GettingToDoctorsAppointmentsDifficultyMap]): void;

  getAdlHelpNecessity(): AdlHelpNecessityMap[keyof AdlHelpNecessityMap];
  setAdlHelpNecessity(value: AdlHelpNecessityMap[keyof AdlHelpNecessityMap]): void;

  getGettingToGroceryStoreDifficulty(): GettingToGroceryStoreDifficultyMap[keyof GettingToGroceryStoreDifficultyMap];
  setGettingToGroceryStoreDifficulty(value: GettingToGroceryStoreDifficultyMap[keyof GettingToGroceryStoreDifficultyMap]): void;

  getSocialIsolation(): SocialIsolationMap[keyof SocialIsolationMap];
  setSocialIsolation(value: SocialIsolationMap[keyof SocialIsolationMap]): void;

  getMealPlanningAssistance(): MealPlanningAssistanceMap[keyof MealPlanningAssistanceMap];
  setMealPlanningAssistance(value: MealPlanningAssistanceMap[keyof MealPlanningAssistanceMap]): void;

  getWeightLossProgramInterest(): WeightLossProgramInterestMap[keyof WeightLossProgramInterestMap];
  setWeightLossProgramInterest(value: WeightLossProgramInterestMap[keyof WeightLossProgramInterestMap]): void;

  getSnapStatus(): SnapStatusMap[keyof SnapStatusMap];
  setSnapStatus(value: SnapStatusMap[keyof SnapStatusMap]): void;

  getFinancialAssistanceStatus(): FinancialAssistanceStatusMap[keyof FinancialAssistanceStatusMap];
  setFinancialAssistanceStatus(value: FinancialAssistanceStatusMap[keyof FinancialAssistanceStatusMap]): void;

  getFoodInsecurity(): FoodInsecurityMap[keyof FoodInsecurityMap];
  setFoodInsecurity(value: FoodInsecurityMap[keyof FoodInsecurityMap]): void;

  getDepressionFrequency(): DepressionFrequencyMap[keyof DepressionFrequencyMap];
  setDepressionFrequency(value: DepressionFrequencyMap[keyof DepressionFrequencyMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PatientSdohQuestionnaire.AsObject;
  static toObject(includeInstance: boolean, msg: PatientSdohQuestionnaire): PatientSdohQuestionnaire.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PatientSdohQuestionnaire, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PatientSdohQuestionnaire;
  static deserializeBinaryFromReader(message: PatientSdohQuestionnaire, reader: jspb.BinaryReader): PatientSdohQuestionnaire;
}

export namespace PatientSdohQuestionnaire {
  export type AsObject = {
    patientId: string,
    patientSdohQuestionnaireId: string,
    questionnaireLanguage: heyrenee_v1_enums_language_pb.LanguageMap[keyof heyrenee_v1_enums_language_pb.LanguageMap],
    dateTimeAnswered?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    ambulatoryStatus: AmbulatoryStatusMap[keyof AmbulatoryStatusMap],
    fallHistory: FallHistoryMap[keyof FallHistoryMap],
    walkingSteadiness: WalkingSteadinessMap[keyof WalkingSteadinessMap],
    walkingSupportNecessity: WalkingSupportNecessityMap[keyof WalkingSupportNecessityMap],
    walkingAssistanceAdvisory: WalkingAssistanceAdvisoryMap[keyof WalkingAssistanceAdvisoryMap],
    walkingAssistanceOwnership: WalkingAssistanceOwnershipMap[keyof WalkingAssistanceOwnershipMap],
    fallingWorry: FallingWorryMap[keyof FallingWorryMap],
    standingUpAssistance: StandingUpAssistanceMap[keyof StandingUpAssistanceMap],
    steppingUpDifficulty: SteppingUpDifficultyMap[keyof SteppingUpDifficultyMap],
    feelingInFeetStatus: FeelingInFeetStatusMap[keyof FeelingInFeetStatusMap],
    feelingLightheadedFromMedsStatus: FeelingLightheadedFromMedsStatusMap[keyof FeelingLightheadedFromMedsStatusMap],
    smokingStatus: SmokingStatusMap[keyof SmokingStatusMap],
    amountSmoked: string,
    quitSmokingStatus: QuitSmokingStatusMap[keyof QuitSmokingStatusMap],
    alcoholDrinkingStatus: AlcoholDrinkingStatusMap[keyof AlcoholDrinkingStatusMap],
    amountAlcoholDrank: string,
    quitDrinkingStatus: QuitDrinkingStatusMap[keyof QuitDrinkingStatusMap],
    schedulingDoctorsAppointmentsDifficulty: SchedulingDoctorsAppointmentsDifficultyMap[keyof SchedulingDoctorsAppointmentsDifficultyMap],
    gettingToDoctorsAppointmentsDifficulty: GettingToDoctorsAppointmentsDifficultyMap[keyof GettingToDoctorsAppointmentsDifficultyMap],
    adlHelpNecessity: AdlHelpNecessityMap[keyof AdlHelpNecessityMap],
    gettingToGroceryStoreDifficulty: GettingToGroceryStoreDifficultyMap[keyof GettingToGroceryStoreDifficultyMap],
    socialIsolation: SocialIsolationMap[keyof SocialIsolationMap],
    mealPlanningAssistance: MealPlanningAssistanceMap[keyof MealPlanningAssistanceMap],
    weightLossProgramInterest: WeightLossProgramInterestMap[keyof WeightLossProgramInterestMap],
    snapStatus: SnapStatusMap[keyof SnapStatusMap],
    financialAssistanceStatus: FinancialAssistanceStatusMap[keyof FinancialAssistanceStatusMap],
    foodInsecurity: FoodInsecurityMap[keyof FoodInsecurityMap],
    depressionFrequency: DepressionFrequencyMap[keyof DepressionFrequencyMap],
  }
}

export interface AmbulatoryStatusMap {
  AMBULATORY_STATUS_UNSPECIFIED: 0;
  AMBULATORY_STATUS_IS_AMBULATORY: 1;
  AMBULATORY_STATUS_NOT_AMBULATORY: 2;
}

export const AmbulatoryStatus: AmbulatoryStatusMap;

export interface FallHistoryMap {
  FALL_HISTORY_UNSPECIFIED: 0;
  FALL_HISTORY_FALLEN_IN_LAST_YEAR: 1;
  FALL_HISTORY_HAS_NOT_FALLEN: 2;
}

export const FallHistory: FallHistoryMap;

export interface WalkingSteadinessMap {
  WALKING_STEADINESS_UNSPECIFIED: 0;
  WALKING_STEADINESS_FEELS_UNSTEADY: 1;
  WALKING_STEADINESS_STEADY: 2;
}

export const WalkingSteadiness: WalkingSteadinessMap;

export interface WalkingSupportNecessityMap {
  WALKING_SUPPORT_NECESSITY_UNSPECIFIED: 0;
  WALKING_SUPPORT_NECESSITY_USES_FURNITURE: 1;
  WALKING_SUPPORT_NECESSITY_NONE: 2;
}

export const WalkingSupportNecessity: WalkingSupportNecessityMap;

export interface WalkingAssistanceAdvisoryMap {
  WALKING_ASSISTANCE_ADVISORY_UNSPECIFIED: 0;
  WALKING_ASSISTANCE_ADVISORY_CANE_OR_WALKER: 1;
  WALKING_ASSISTANCE_ADVISORY_NONE: 2;
}

export const WalkingAssistanceAdvisory: WalkingAssistanceAdvisoryMap;

export interface WalkingAssistanceOwnershipMap {
  WALKING_ASSISTANCE_OWNERSHIP_UNSPECIFIED: 0;
  WALKING_ASSISTANCE_OWNERSHIP_CANE_OR_WALKER: 1;
  WALKING_ASSISTANCE_OWNERSHIP_NONE: 2;
}

export const WalkingAssistanceOwnership: WalkingAssistanceOwnershipMap;

export interface FallingWorryMap {
  FALLING_WORRY_UNSPECIFIED: 0;
  FALLING_WORRY_IS_WORRIED: 1;
  FALLING_WORRY_NOT_WORRIED: 2;
}

export const FallingWorry: FallingWorryMap;

export interface StandingUpAssistanceMap {
  STANDING_UP_ASSISTANCE_UNSPECIFIED: 0;
  STANDING_UP_ASSISTANCE_PUSHES: 1;
  STANDING_UP_ASSISTANCE_NONE: 2;
}

export const StandingUpAssistance: StandingUpAssistanceMap;

export interface SteppingUpDifficultyMap {
  STEPPING_UP_DIFFICULTY_UNSPECIFIED: 0;
  STEPPING_UP_DIFFICULTY_HAS_DIFFICULTY: 1;
  STEPPING_UP_DIFFICULTY_NONE: 2;
}

export const SteppingUpDifficulty: SteppingUpDifficultyMap;

export interface FeelingInFeetStatusMap {
  FEELING_IN_FEET_STATUS_UNSPECIFIED: 0;
  FEELING_IN_FEET_STATUS_LOST_SOME: 1;
  FEELING_IN_FEET_STATUS_HAS_ALL: 2;
}

export const FeelingInFeetStatus: FeelingInFeetStatusMap;

export interface FeelingLightheadedFromMedsStatusMap {
  FEELING_LIGHTHEADED_FROM_MEDS_STATUS_UNSPECIFIED: 0;
  FEELING_LIGHTHEADED_FROM_MEDS_STATUS_AT_LEAST_ONE: 1;
  FEELING_LIGHTHEADED_FROM_MEDS_STATUS_NO_MEDS: 2;
}

export const FeelingLightheadedFromMedsStatus: FeelingLightheadedFromMedsStatusMap;

export interface SmokingStatusMap {
  SMOKING_STATUS_UNSPECIFIED: 0;
  SMOKING_STATUS_SMOKES: 1;
  SMOKING_STATUS_DOES_NOT_SMOKE: 2;
}

export const SmokingStatus: SmokingStatusMap;

export interface QuitSmokingStatusMap {
  QUIT_SMOKING_STATUS_UNSPECIFIED: 0;
  QUIT_SMOKING_STATUS_WANTS_TO_QUIT: 1;
  QUIT_SMOKING_STATUS_DOES_NOT_WANT_TO_QUIT: 2;
}

export const QuitSmokingStatus: QuitSmokingStatusMap;

export interface AlcoholDrinkingStatusMap {
  ALCOHOL_DRINKING_STATUS_UNSPECIFIED: 0;
  ALCOHOL_DRINKING_STATUS_DRINKS: 1;
  ALCOHOL_DRINKING_STATUS_DOES_NOT_DRINK: 2;
}

export const AlcoholDrinkingStatus: AlcoholDrinkingStatusMap;

export interface QuitDrinkingStatusMap {
  QUIT_DRINKING_STATUS_UNSPECIFIED: 0;
  QUIT_DRINKING_STATUS_WANTS_TO_QUIT: 1;
  QUIT_DRINKING_STATUS_DOES_NOT_WANT_TO_QUIT: 2;
}

export const QuitDrinkingStatus: QuitDrinkingStatusMap;

export interface SchedulingDoctorsAppointmentsDifficultyMap {
  SCHEDULING_DOCTORS_APPOINTMENTS_DIFFICULTY_UNSPECIFIED: 0;
  SCHEDULING_DOCTORS_APPOINTMENTS_DIFFICULTY_HAS_DIFFICULTY: 1;
  SCHEDULING_DOCTORS_APPOINTMENTS_DIFFICULTY_NO_DIFFICULTY: 2;
}

export const SchedulingDoctorsAppointmentsDifficulty: SchedulingDoctorsAppointmentsDifficultyMap;

export interface GettingToDoctorsAppointmentsDifficultyMap {
  GETTING_TO_DOCTORS_APPOINTMENTS_DIFFICULTY_UNSPECIFIED: 0;
  GETTING_TO_DOCTORS_APPOINTMENTS_DIFFICULTY_TROUBLE_FINDING_RIDE: 1;
  GETTING_TO_DOCTORS_APPOINTMENTS_DIFFICULTY_NO_DIFFICULTY: 2;
}

export const GettingToDoctorsAppointmentsDifficulty: GettingToDoctorsAppointmentsDifficultyMap;

export interface AdlHelpNecessityMap {
  ADL_HELP_NECESSITY_UNSPECIFIED: 0;
  ADL_HELP_NECESSITY_NEEDED: 1;
  ADL_HELP_NECESSITY_NOT_NEEDED: 2;
}

export const AdlHelpNecessity: AdlHelpNecessityMap;

export interface GettingToGroceryStoreDifficultyMap {
  GETTING_TO_GROCERY_STORE_DIFFICULTY_UNSPECIFIED: 0;
  GETTING_TO_GROCERY_STORE_DIFFICULTY_UNABLE_ALONE: 1;
  GETTING_TO_GROCERY_STORE_DIFFICULTY_ABLE_ALONE: 2;
}

export const GettingToGroceryStoreDifficulty: GettingToGroceryStoreDifficultyMap;

export interface SocialIsolationMap {
  SOCIAL_ISOLATION_UNSPECIFIED: 0;
  SOCIAL_ISOLATION_FEELS_ISOLATED: 1;
  SOCIAL_ISOLATION_DOES_NOT_FEEL_ISOLATED: 2;
}

export const SocialIsolation: SocialIsolationMap;

export interface MealPlanningAssistanceMap {
  MEAL_PLANNING_ASSISTANCE_UNSPECIFIED: 0;
  MEAL_PLANNING_ASSISTANCE_WANTS: 1;
  MEAL_PLANNING_ASSISTANCE_DOES_NOT_WANT: 2;
}

export const MealPlanningAssistance: MealPlanningAssistanceMap;

export interface WeightLossProgramInterestMap {
  WEIGHT_LOSS_PROGRAM_INTEREST_UNSPECIFIED: 0;
  WEIGHT_LOSS_PROGRAM_INTEREST_INTERESTED: 1;
  WEIGHT_LOSS_PROGRAM_INTEREST_NOT_INTERESTED: 2;
}

export const WeightLossProgramInterest: WeightLossProgramInterestMap;

export interface SnapStatusMap {
  SNAP_STATUS_UNSPECIFIED: 0;
  SNAP_STATUS_ON_SNAP: 1;
  SNAP_STATUS_NOT_ON_SNAP: 2;
}

export const SnapStatus: SnapStatusMap;

export interface FinancialAssistanceStatusMap {
  FINANCIAL_ASSISTANCE_STATUS_UNSPECIFIED: 0;
  FINANCIAL_ASSISTANCE_STATUS_RECEIVES_FINANCIAL_ASSISTANCE: 1;
  FINANCIAL_ASSISTANCE_STATUS_NO_FINANCIAL_ASSISTANCE: 2;
}

export const FinancialAssistanceStatus: FinancialAssistanceStatusMap;

export interface FoodInsecurityMap {
  FOOD_INSECURITY_UNSPECIFIED: 0;
  FOOD_INSECURITY_INSECURE: 1;
  FOOD_INSECURITY_SECURE: 2;
}

export const FoodInsecurity: FoodInsecurityMap;

export interface DepressionFrequencyMap {
  DEPRESSION_FREQUENCY_UNSPECIFIED: 0;
  DEPRESSION_FREQUENCY_FREQUENT: 1;
  DEPRESSION_FREQUENCY_INFREQUENT: 2;
}

export const DepressionFrequency: DepressionFrequencyMap;

