import { getOutcomeColor, getOutcomeOptions } from "@/app/(app)/apis/[apiId]/_overview/utils";
import { FilterCheckbox } from "@/components/logs/checkbox/filter-checkbox";
import { useFilters } from "../../../../hooks/use-filters";

type OutcomesFilterProps = {
  // onDrawerClose is a callback function that will be called when applying the filter
  onDrawerClose?: () => void;
};

export const OutcomesFilter = ({ onDrawerClose }: OutcomesFilterProps) => {
  const { filters, updateFilters } = useFilters();

  // Get the options from our centralized utility
  const options = getOutcomeOptions();

  return (
    <FilterCheckbox
      options={options}
      filterField="outcomes"
      checkPath="outcome"
      renderOptionContent={(checkbox) => (
        <>
          <div className={`size-2 ${checkbox.color} rounded-[2px]`} />
          <span className="text-accent-12 text-xs">{checkbox.label}</span>
        </>
      )}
      createFilterValue={(option) => ({
        value: option.outcome,
        metadata: {
          colorClass: getOutcomeColor(option.outcome),
        },
      })}
      filters={filters}
      updateFilters={updateFilters}
      onDrawerClose={onDrawerClose}
    />
  );
};
